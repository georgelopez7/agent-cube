import { cn } from "cnfast";
import { Wifi, WifiOff } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import {
  CubeAgentExitEventHandler,
  CubeAgentReasoningEventHandler,
  CubeAgentStoppedEventHandler,
  CubeAgentTimeoutEventHandler,
  CubeCompletedEventHandler,
  CubeRotatedEventHandler,
  CubeUsageEventHandler,
  type IEvent,
  WebSocketEventType,
} from "./websocket-events";

interface WebsocketProps {
  debug?: boolean;
}

const Websocket = ({ debug = false }: WebsocketProps) => {
  const socket = useRef<WebSocket | null>(null);
  const [connected, setConnected] = useState(false);

  const endpoint =
    import.meta.env.VITE_WEBSOCKET_URL ?? "ws://localhost:8080/api/v1/ws";

  useEffect(() => {
    socket.current = new WebSocket(endpoint);

    socket.current.onopen = () => {
      setConnected(true);
    };

    socket.current.onmessage = async (event) => {
      const ev = JSON.parse(event.data) as IEvent;

      if (debug) console.log(ev);

      switch (ev.type) {
        case WebSocketEventType.CubeRotated:
          await CubeRotatedEventHandler(ev);
          break;
        case WebSocketEventType.CubeAgentReasoning:
          await CubeAgentReasoningEventHandler(ev);
          break;
        case WebSocketEventType.CubeAgentTimeout:
          await CubeAgentTimeoutEventHandler(ev);
          break;
        case WebSocketEventType.CubeAgentStopped:
          await CubeAgentStoppedEventHandler(ev);
          break;
        case WebSocketEventType.CubeAgentExit:
          await CubeAgentExitEventHandler(ev);
          break;
        case WebSocketEventType.CubeCompleted:
          await CubeCompletedEventHandler(ev);
          break;
        case WebSocketEventType.CubeUsage:
          await CubeUsageEventHandler(ev);
          break;
        default:
          console.warn("unknown websocket event", ev);
          break;
      }
    };

    socket.current.onclose = async () => {
      setConnected(false);
    };

    return () => {
      if (!socket.current) return;
      socket.current.close();
    };
  }, []);

  if (!debug) {
    return null;
  }

  // DEBUG COMPONENT

  const config = {
    connected: {
      icon: Wifi,
      label: "Connected",
      color: "bg-green-600",
    },
    disconnected: {
      icon: WifiOff,
      label: "Disconnected",
      color: "bg-red-600",
    },
  };

  const status = connected ? config.connected : config.disconnected;
  const Icon = status.icon;

  return (
    <div
      className={cn(
        "fixed top-0 left-0 right-0 z-50 flex items-center justify-center gap-2 px-4 py-2 text-sm font-medium text-white",
        status.color,
      )}
    >
      <Icon className="h-4 w-4" />
      {status.label}
    </div>
  );
};

export default Websocket;
