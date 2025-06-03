from smolagents import Tool, CodeAgent, InferenceClientModel, tool
from dotenv import load_dotenv

load_dotenv()

# Method 1: Using @tool decorator
@tool
def superhero_power_analyzer(hero_name: str) -> str:
    """
    Analyzes a superhero's powers and provides a detailed breakdown of their abilities.

    Args:
        hero_name: Name of the superhero to analyze.
    """
    powers_database = {
        "Batman": {
            "powers": ["Peak Human Condition", "Master Detective", "Martial Arts Expert"],
            "weaknesses": ["No Superpowers", "Human Limitations"],
            "equipment": ["Batsuit", "Utility Belt", "Batmobile"]
        },
        "Superman": {
            "powers": ["Super Strength", "Flight", "Heat Vision", "X-Ray Vision"],
            "weaknesses": ["Kryptonite", "Magic"],
            "equipment": ["Superman Suit"]
        },
        "Wonder Woman": {
            "powers": ["Superhuman Strength", "Flight", "Lasso of Truth"],
            "weaknesses": ["Bound by Truth", "Vulnerable to Piercing Weapons"],
            "equipment": ["Lasso of Truth", "Bracelets of Submission", "Tiara"]
        }
    }
    
    hero = powers_database.get(hero_name, None)
    if not hero:
        return f"Sorry, I don't have information about {hero_name}."
    
    return f"""
    Power Analysis for {hero_name}:
    Powers: {', '.join(hero['powers'])}
    Weaknesses: {', '.join(hero['weaknesses'])}
    Equipment: {', '.join(hero['equipment'])}
    """

# Method 2: Using custom Tool class
class SuperheroTeamBuilder(Tool):
    name = "superhero_team_builder"
    description = """
    This tool helps create balanced superhero teams based on mission requirements.
    It suggests team compositions with complementary powers and skills."""

    inputs = {
        "mission_type": {
            "type": "string",
            "description": "The type of mission (e.g., 'stealth', 'combat', 'rescue', 'investigation').",
        }
    }

    output_type = "string"

    def forward(self, mission_type: str):
        team_compositions = {
            "stealth": {
                "leader": "Batman",
                "members": ["Catwoman", "Nightwing", "Green Arrow"],
                "strategy": "Covert operations and infiltration"
            },
            "combat": {
                "leader": "Superman",
                "members": ["Wonder Woman", "Martian Manhunter", "Shazam"],
                "strategy": "Direct confrontation and heavy combat"
            },
            "rescue": {
                "leader": "Flash",
                "members": ["Superman", "Green Lantern", "Cyborg"],
                "strategy": "Rapid response and civilian protection"
            },
            "investigation": {
                "leader": "Batman",
                "members": ["Constantine", "Zatanna", "Detective Chimp"],
                "strategy": "Mystical and scientific investigation"
            }
        }

        team = team_compositions.get(mission_type.lower())
        if not team:
            return "Mission type not found. Try 'stealth', 'combat', 'rescue', or 'investigation'."

        return f"""
        Recommended Team for {mission_type.title()} Mission:
        Team Leader: {team['leader']}
        Team Members: {', '.join(team['members'])}
        Strategy: {team['strategy']}
        """

# Create instances of both tools
power_analyzer = superhero_power_analyzer
team_builder = SuperheroTeamBuilder()

# Create agent with multiple tools and specify the provider
model = InferenceClientModel(provider="hf-inference")
agent = CodeAgent(
    tools=[power_analyzer, team_builder],
    model=model
)

# Example queries to demonstrate both tools
queries = [
    "Can you suggest a team for a stealth mission?",
    "I need a team for a rescue mission, who should I include?"
]

# Run the agent with different queries
for query in queries:
    print(f"\nQuery: {query}")
    result = agent.run(query)
    print(f"Result: {result}")