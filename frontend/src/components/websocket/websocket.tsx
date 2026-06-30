import { useEffect, useRef } from "react";
import { RotationEventHandler, type IEvent } from "./websocket-events";

const Websocket = () => {
  const socket = useRef<WebSocket | null>(null);

  useEffect(() => {
    socket.current = new WebSocket(process.env.WEBSOCKET_URL!);

    socket.current.onopen = () => {};

    socket.current.onmessage = async (event) => {
      const payload = JSON.parse(event.data) as IEvent;

      switch (payload.event_type) {
        case "cube.rotation":
          await RotationEventHandler(payload);
          break;
        // case "cube.solved":
        //   await SolvedEventHandler(payload);
        //   break;
        // case "agent.message":
        //   await AgentMessageEventHandler(payload);
        //   break;
        default:
          break;
      }
    };

    socket.current.onclose = async () => {};

    return () => {
      if (!socket.current) return;
      socket.current.close();
    };
  }, []);

  return <></>;
};

export default Websocket;
