#!/usr/bin/env node
import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { CallToolRequestSchema, ListToolsRequestSchema, ToolSchema, } from "@modelcontextprotocol/sdk/types.js";
import fs from "fs/promises";
import path from "path";
import { z } from "zod";
import { zodToJsonSchema } from "zod-to-json-schema";
// Create a custom logger that writes to a file instead of console
const logFile = path.join(process.cwd(), 'mcp-server.log');
const logger = {
    log: async (message) => {
        try {
            const formattedMessage = typeof message === 'string'
                ? message
                : JSON.stringify(message, null, 2);
            await fs.appendFile(logFile, formattedMessage + '\n', 'utf8');
        }
        catch (e) {
            // Silently fail if we can't write to the log file
        }
    }
};
// Command line argument parsing
const args = process.argv.slice(2);
if (args.length === 0) {
    logger.log("Usage: mcp-server-filesystem <allowed-directory> [additional-directories...]");
    process.exit(1);
}
// Normalize all paths consistently
function normalizePath(p) {
    return path.normalize(p);
}
// Store allowed directories in normalized form
const allowedDirectories = args.map(dir => normalizePath(path.resolve(dir)));
// Get the routines directory from args
const ROUTINES_DIR = allowedDirectories[0]; // Use the first provided directory
logger.log(`Setting ROUTINES_DIR to: ${ROUTINES_DIR}`);
// Validate that all directories exist and are accessible
await Promise.all(args.map(async (dir) => {
    try {
        const stats = await fs.stat(dir);
        if (!stats.isDirectory()) {
            logger.log(`Error: ${dir} is not a directory`);
            process.exit(1);
        }
        logger.log(`Successfully validated directory: ${dir}`);
    }
    catch (error) {
        logger.log(`Error accessing directory ${dir}: ${error}`);
        process.exit(1);
    }
}));
// Security utilities
async function validatePath(requestedPath) {
    logger.log(`Validating path: ${requestedPath}`);
    const absolute = path.isAbsolute(requestedPath)
        ? path.resolve(requestedPath)
        : path.resolve(ROUTINES_DIR, requestedPath);
    logger.log(`Resolved absolute path: ${absolute}`);
    const normalizedRequested = normalizePath(absolute);
    logger.log(`Normalized path: ${normalizedRequested}`);
    // Check if path is within allowed directories
    const isAllowed = allowedDirectories.some(dir => {
        const result = normalizedRequested.startsWith(dir);
        logger.log(`Checking against allowed directory ${dir}: ${result}`);
        return result;
    });
    if (!isAllowed) {
        logger.log(`❌ Access denied - path outside allowed directories`);
        logger.log(`Path: ${absolute}`);
        logger.log(`Allowed directories: ${allowedDirectories.join(', ')}`);
        throw new Error(`Access denied - path outside allowed directories: ${absolute} not in ${allowedDirectories.join(', ')}`);
    }
    logger.log(`✅ Path validated successfully: ${absolute}`);
    return absolute;
}
// Schema definitions
const RoutineStepSchema = z.object({
    name: z.string(),
    count: z.string(),
    description: z.string()
});
const DanceRoutineSchema = z.object({
    name: z.string(),
    description: z.string(),
    difficulty: z.string(),
    duration: z.string(),
    style: z.string(),
    creator: z.string(),
    year: z.string(),
    steps: z.array(RoutineStepSchema)
});
const GetRoutineArgsSchema = z.object({
    name: z.string().describe('Name of the routine to retrieve')
});
const ListRoutinesArgsSchema = z.object({
    style: z.string().optional().describe('Optional filter by dance style'),
    difficulty: z.string().optional().describe('Optional filter by difficulty level')
});
const ToolInputSchema = ToolSchema.shape.inputSchema;
// Server setup
const server = new Server({
    name: "dance-routine-server",
    version: "1.0.0",
}, {
    capabilities: {
        tools: {},
    },
});
logger.log('=== SERVER STARTING ===');
// Tool handlers
server.setRequestHandler(ListToolsRequestSchema, async () => {
    logger.log('🔍 LIST TOOLS REQUEST RECEIVED');
    logger.log('Available tools: get_routine, list_routines');
    return {
        tools: [
            {
                name: "get_routine",
                description: "Retrieve a specific dance routine by name. Returns the complete routine information including steps, difficulty, and other metadata.",
                inputSchema: zodToJsonSchema(GetRoutineArgsSchema),
            },
            {
                name: "list_routines",
                description: "List all available dance routines. Can be filtered by style and difficulty level.",
                inputSchema: zodToJsonSchema(ListRoutinesArgsSchema),
            }
        ],
    };
});
server.setRequestHandler(CallToolRequestSchema, async (request) => {
    logger.log(`🛠 CALL TOOL REQUEST: ${request.params.name}`);
    logger.log(`Arguments: ${JSON.stringify(request.params.arguments, null, 2)}`);
    try {
        const { name, arguments: args } = request.params;
        switch (name) {
            case "get_routine": {
                logger.log('📖 Getting routine...');
                const parsed = GetRoutineArgsSchema.safeParse(args);
                if (!parsed.success) {
                    throw new Error(`Invalid arguments for get_routine: ${parsed.error}`);
                }
                // Use absolute path for routines
                const routinePath = path.join(ROUTINES_DIR, `${parsed.data.name.toLowerCase().replace(/\s+/g, '-')}-full-routine.json`);
                logger.log(`Looking for routine at: ${routinePath}`);
                const validPath = await validatePath(routinePath);
                try {
                    const content = await fs.readFile(validPath, "utf-8");
                    const routine = JSON.parse(content);
                    const validatedRoutine = DanceRoutineSchema.parse(routine);
                    return {
                        content: [{ type: "text", text: JSON.stringify(validatedRoutine, null, 2) }],
                    };
                }
                catch (error) {
                    throw new Error(`Failed to read routine ${parsed.data.name}: ${error}`);
                }
            }
            case "list_routines": {
                logger.log('📋 Listing routines...');
                const parsed = ListRoutinesArgsSchema.safeParse(args);
                if (!parsed.success) {
                    logger.log('❌ Invalid arguments:');
                    logger.log(parsed.error);
                    throw new Error(`Invalid arguments for list_routines: ${parsed.error}`);
                }
                // Use absolute path for routines
                const validPath = await validatePath(ROUTINES_DIR);
                logger.log('📁 Reading directory:');
                logger.log(validPath);
                const files = await fs.readdir(validPath);
                logger.log('📄 Found files:');
                logger.log(files);
                const routines = await Promise.all(files
                    .filter(file => file.endsWith('.json'))
                    .map(async (file) => {
                    const content = await fs.readFile(path.join(ROUTINES_DIR, file), 'utf-8');
                    return JSON.parse(content);
                }));
                const filteredRoutines = routines.filter(routine => {
                    if (parsed.data.style && routine.style !== parsed.data.style)
                        return false;
                    if (parsed.data.difficulty && routine.difficulty !== parsed.data.difficulty)
                        return false;
                    return true;
                });
                return {
                    content: [{
                            type: "text",
                            text: JSON.stringify(filteredRoutines.map(r => ({
                                name: r.name,
                                style: r.style,
                                difficulty: r.difficulty,
                                duration: r.duration
                            })), null, 2)
                        }],
                };
            }
            default:
                throw new Error(`Unknown tool: ${name}`);
        }
    }
    catch (error) {
        logger.log('❌ Error:');
        logger.log(error);
        const errorMessage = error instanceof Error ? error.message : String(error);
        return {
            content: [{ type: "text", text: `Error: ${errorMessage}` }],
            isError: true,
        };
    }
});
// Start server
async function runServer() {
    logger.log('🚀 Starting Dance Routine Server...');
    // Use process.stdout and process.stdin directly for MCP communication
    const transport = new StdioServerTransport();
    // Set up error handling
    process.on('uncaughtException', (error) => {
        logger.log(`💥 Uncaught Exception: ${error.message}`);
        logger.log(`${error.stack}`);
    });
    process.on('unhandledRejection', (reason, promise) => {
        logger.log('🔥 Unhandled Rejection at:');
        logger.log(`Reason: ${reason}`);
    });
    // Connect the server
    await server.connect(transport);
    logger.log('✅ Dance Routine Server running on stdio');
    logger.log(`📁 Allowed directories: ${allowedDirectories.join(', ')}`);
    // Keep the process alive
    process.stdin.resume();
}
runServer().catch((error) => {
    logger.log(`💥 Fatal error running server: ${error.message}`);
    logger.log(`${error.stack}`);
    process.exit(1);
});
