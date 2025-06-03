from smolagents import CodeAgent, DuckDuckGoSearchTool, InferenceClientModel
from dotenv import load_dotenv

load_dotenv()

agent = CodeAgent(
    tools=[DuckDuckGoSearchTool()],
    model=InferenceClientModel(provider="hf-inference")
    )

agent.run("Find the latest news about artificial intelligence in law, give me top 3 links.")