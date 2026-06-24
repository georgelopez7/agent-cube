from langchain_core.messages import AIMessage, HumanMessage
from langfuse.langchain import CallbackHandler

from internal.pkg.agent.agent import NewRubiksCubeAgent
from internal.pkg.agent.prompts import NewStarterPrompt


class RubiksCubeService:
    def __init__(self) -> None:
        pass

    def RunRubiksCubeAgent(self, id: str, model: str) -> str:
        """
        Runs the Rubik's Cube Agent with the given ID and model.
        """

        langfuse_handler = CallbackHandler()

        agent = NewRubiksCubeAgent(model)

        prompt = NewStarterPrompt(id)
        messages = {"messages": [HumanMessage(content=prompt)]}

        final_content = ""
        for chunk in agent.stream(
            messages, stream_mode="values", config={"callbacks": [langfuse_handler]}
        ):
            latest_message = chunk["messages"][-1]

            # TOOL CALLS
            if hasattr(latest_message, "tool_calls") and latest_message.tool_calls:
                message = f"[TOOL] - {[tc['name'] for tc in latest_message.tool_calls]}"
                print(message, flush=True)

            # AGENT MESSAGES
            if isinstance(latest_message, AIMessage) and latest_message.content:
                message = f"[AGENT] - {latest_message.content}"
                print(message, flush=True)

            # FINAL MESSAGE
            if isinstance(latest_message, AIMessage):
                final_content = latest_message.content

        return final_content
