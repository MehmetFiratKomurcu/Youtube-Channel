from smolagents import CodeAgent, InferenceClientModel, tool
from dotenv import load_dotenv
import gradio as gr
from PIL import Image
import io
import base64
import requests
import os
import json
import re

load_dotenv()

# Agent 1: Artist Agent - Generates images from text prompts
@tool
def generate_image(prompt: str) -> str:
    """
    Generates an image from a text prompt using Stable Diffusion.
    
    Args:
        prompt: The text description of the image to generate.
    """
    # Hugging Face API endpoint for Stable Diffusion
    API_URL = "https://api-inference.huggingface.co/models/stabilityai/stable-diffusion-xl-base-1.0"
    headers = {"Authorization": f"Bearer {os.getenv('HF_TOKEN')}"}

    # Make the API request
    response = requests.post(API_URL, headers=headers, json={"inputs": prompt})
    
    if response.status_code != 200:
        return "Error generating image"
    
    # Convert the response to an image
    image = Image.open(io.BytesIO(response.content))
    
    # Save the image temporarily and return the path
    image_path = "generated_image.png"
    image.save(image_path)
    return image_path

# Agent 2: Excited Child Agent - Describes images in a playful way
@tool
def describe_image_excitedly(image_path: str) -> str:
    """
    Describes an image in an excited, child-like manner.
    
    Args:
        image_path: Path to the image to describe.
    """
    try:
        # Load the image and convert to base64
        with open(image_path, "rb") as image_file:
            base64_image = base64.b64encode(image_file.read()).decode('utf-8')
        
        # Create a prompt for the LLM
        prompt = f"""
        You are a 5-year-old child who just saw an amazing picture! 
        Look at this image and describe it with the excitement and wonder of a child.
        Use simple words, exclamation marks, and show your amazement!
        Be playful and use your imagination!
        
        Image: [IMAGE]
        """
        
        # Use the LLM to generate the description
        model = InferenceClientModel(provider="hf-inference")
        response = model.generate(prompt, images=[base64_image])
        
        # Debug logging
        print(f"LLM Response Type: {type(response)}")
        print(f"LLM Response: {response}")
        
        # Handle the response based on its type
        if isinstance(response, dict):
            return response.get("generated_text", str(response))
        elif isinstance(response, list):
            return response[0] if response else "No description generated"
        elif hasattr(response, "generated_text"):
            return response.generated_text
        elif hasattr(response, "text"):
            return response.text
        elif hasattr(response, "content"):
            return response.content
        else:
            return str(response)
        
    except Exception as e:
        print(f"Error in describe_image_excitedly: {str(e)}")
        return f"Error describing image: {str(e)}"

# Create the specialized agents
artist_agent = CodeAgent(
    tools=[generate_image],
    model=InferenceClientModel(provider="hf-inference"),
    name="artist",
    description="An agent that generates images from text descriptions using Stable Diffusion."
)

child_agent = CodeAgent(
    tools=[describe_image_excitedly],
    model=InferenceClientModel(provider="hf-inference"),
    name="child",
    description="An agent that describes images in an excited, child-like manner."
)

# Create the manager agent that coordinates the workflow
manager_agent = CodeAgent(
    model=InferenceClientModel(provider="hf-inference"),
    tools=[generate_image, describe_image_excitedly],
    managed_agents=[artist_agent, child_agent],
    planning_interval=5,
    verbosity_level=2,
    max_steps=15
)

def extract_image_path(text):
    """Extract the image path from the manager agent's response."""
    # First try to find a JSON-like structure
    json_match = re.search(r'\{[^}]+\}', text)
    if json_match:
        try:
            result = json.loads(json_match.group())
            if isinstance(result, dict) and "image_path" in result:
                return result["image_path"]
        except json.JSONDecodeError:
            pass
    
    # If no JSON found, look for common image file patterns
    image_pattern = r'[\w\-\.]+\.(png|jpg|jpeg|gif)'
    match = re.search(image_pattern, text)
    if match:
        return match.group()
    
    # If no image path found, return the default path
    return "generated_image.png"

def extract_description(text):
    """Extract the description from the manager agent's response."""
    try:
        # First try to parse as JSON
        if isinstance(text, str):
            result = json.loads(text)
        else:
            result = text
            
        # If it's a dictionary and has a description key, return only that
        if isinstance(result, dict) and "description" in result:
            return result["description"]
            
        # If it's a dictionary but no description key, return the whole text
        return str(result)
    except (json.JSONDecodeError, AttributeError):
        # If not JSON, return the text as is
        return str(text)

# Gradio interface
def process_prompt(prompt):
    # The manager agent coordinates the workflow
    print("🤖 Manager Agent is coordinating the workflow...")
    
    # Define the workflow for the manager agent
    workflow = f"""
    Task: Generate an image and get an excited child's description of it.
    
    Steps:
    1. First, use the artist agent to create an image:
       image_path = artist(task="Generate an image of {prompt}")
       print(f"Generated image saved at: {{image_path}}")
    
    2. Then, use the child agent to describe the image:
       description = child(task="Describe this image in an excited, child-like manner: {{image_path}}")
       print(f"Child's description: {{description}}")
    
    3. Finally, return the results:
       final_answer(
           image_path=image_path,
           description=description
       )
    """
    
    # Run the managed workflow
    result = manager_agent.run(workflow)
    print(f"Manager Agent Response: {result}")
    
    # Extract image path and description from the result
    image_path = extract_image_path(str(result))
    description = extract_description(result)
    
    print(f"Extracted Image Path: {image_path}")
    print(f"Extracted Description: {description}")
    
    return image_path, description

# Create the Gradio interface
iface = gr.Interface(
    fn=process_prompt,
    inputs=gr.Textbox(label="Enter your image description"),
    outputs=[
        gr.Image(label="Generated Image"),
        gr.Textbox(label="Child's Excited Description")
    ],
    title="🎨 Managed Multi-Agent System",
    description="Watch how the Manager Agent coordinates the Artist and Child agents to create and describe images!"
)

# Launch the interface
if __name__ == "__main__":
    iface.launch()
