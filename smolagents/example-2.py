from smolagents import CodeAgent, InferenceClientModel, tool
from dotenv import load_dotenv

load_dotenv()

# Creating a custom tool using the @tool decorator
@tool
def superhero_training_center(query: str) -> str:
    """
    This tool returns the best superhero training centers in Gotham and their features.

    Args:
        query: Search term for training centers.
    """
    # Example training centers and their features
    training_centers = {
        "Wayne Enterprises Training Facility": {
            "rating": 4.9,
            "specialties": ["Combat Training", "Detective Skills", "Gadget Mastery"],
            "notable_instructors": ["Batman", "Alfred Pennyworth"]
        },
        "Themyscira Academy": {
            "rating": 4.8,
            "specialties": ["Amazonian Combat", "Ancient Weapons", "Leadership"],
            "notable_instructors": ["Wonder Woman", "Hippolyta"]
        },
        "Atlantean Training Grounds": {
            "rating": 4.7,
            "specialties": ["Underwater Combat", "Ocean Magic", "Marine Biology"],
            "notable_instructors": ["Aquaman", "Mera"]
        }
    }

    # Find the highest rated center
    best_center = max(training_centers.items(), key=lambda x: x[1]["rating"])
    
    # Format detailed information
    result = f"""
    Best Training Center: {best_center[0]}
    Rating: {best_center[1]['rating']}
    Specialties: {', '.join(best_center[1]['specialties'])}
    Notable Instructors: {', '.join(best_center[1]['notable_instructors'])}
    """
    
    return result

# Create agent and add the tool
agent = CodeAgent(
    tools=[superhero_training_center],
    model=InferenceClientModel(provider="hf-inference")
)

# Run the agent
result = agent.run(
    "Can you find the best superhero training center in Gotham and its features?"
)

print(result)