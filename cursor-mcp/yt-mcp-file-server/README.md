# Dance Routine File Server

This server provides access to dance routine files stored in the `routines` directory. It supports reading individual routines and listing all available routines with optional filtering.

## Available Tools

### 1. Get Routine (`get_routine`)
Retrieves a specific dance routine by name. Returns the complete routine information including steps, difficulty, and other metadata.

Example usage:
```json
{
  "name": "Shim Sham"
}
```

### 2. List Routines (`list_routines`)
Lists all available dance routines. Can be filtered by style and difficulty level.

Example usage:
```json
{
  "style": "jazz",
  "difficulty": "intermediate"
}
```

## Routine File Format

Each routine is stored as a JSON file in the `routines` directory with the following structure:

```json
{
  "name": "Routine Name",
  "description": "Routine description",
  "difficulty": "beginner|intermediate|advanced",
  "duration": "duration in bars or minutes",
  "style": "dance style",
  "creator": "creator name",
  "year": "year created",
  "steps": [
    {
      "name": "Step Name",
      "count": "number of counts",
      "description": "step description"
    }
  ]
}
```

## File Naming Convention

Routine files should be named using the following format:
`{routine-name-in-lowercase}-full-routine.json`

Example: `shim-sham-full-routine.json`

## Installation

1. Install the required dependencies:
```bash
npm install
```

2. Start the server:
```bash
npm start
```

## Adding New Routines

To add a new routine:

1. Create a new JSON file in the `routines` directory following the naming convention
2. Follow the routine file format structure
3. The server will automatically detect and serve the new routine

## Inspector Command

npx @modelcontextprotocol/inspector node dist/index.js