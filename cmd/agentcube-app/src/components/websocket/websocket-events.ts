import {
  type CubeRotation,
  RubiksCubeStatus,
  type TokenUsage,
} from "#/domain/rubiks-cube";
import { useRubiksCubeStore } from "#/stores/rubiks-cube-store";

export enum WebSocketEventType {
  CubeRotated = "cube.rotated",
  CubeAgentReasoning = "cube.agent.reasoning",
  CubeAgentTimeout = "cube.agent.timeout",
  CubeCompleted = "cube.completed",
  CubeUsage = "cube.usage",
}

export type CubeRotatedEvent = {
  type: WebSocketEventType.CubeRotated;
  payload: {
    cube_id: string;
    rotation: CubeRotation;
  };
};

// cube.rotated
export const CubeRotatedEventHandler = async (event: CubeRotatedEvent) => {
  const { records, AddRotation } = useRubiksCubeStore.getState();

  const { cube_id, rotation } = event.payload;

  const cube = records.get(cube_id);

  if (!cube) return;

  AddRotation(cube_id, rotation);
  await cube.ref.rotate(rotation.rotation);
};

export type CubeAgentTimeoutEvent = {
  type: WebSocketEventType.CubeAgentTimeout;
  payload: {
    cube_id: string;
  };
};

export type CubeAgentReasoningEvent = {
  type: WebSocketEventType.CubeAgentReasoning;
  payload: {
    cube_id: string;
    text: string;
  };
};

// cube.agent.reasoning
export const CubeAgentReasoningEventHandler = async (
  event: CubeAgentReasoningEvent,
) => {
  const { AppendReasoning } = useRubiksCubeStore.getState();

  AppendReasoning(event.payload.cube_id, event.payload.text);
};

// cube.agent.timeout
export const CubeAgentTimeoutEventHandler = async (
  event: CubeAgentTimeoutEvent,
) => {
  const { SetStatus } = useRubiksCubeStore.getState();

  SetStatus(event.payload.cube_id, RubiksCubeStatus.TimedOut);
};

export type CubeCompletedEvent = {
  type: WebSocketEventType.CubeCompleted;
  payload: {
    cube_id: string;
  };
};

// cube.completed
export const CubeCompletedEventHandler = async (event: CubeCompletedEvent) => {
  const { SetStatus } = useRubiksCubeStore.getState();

  SetStatus(event.payload.cube_id, RubiksCubeStatus.Completed);
};

export type CubeUsageEvent = {
  type: WebSocketEventType.CubeUsage;
  payload: {
    cube_id: string;
    usage: TokenUsage;
  };
};

// cube.usage
export const CubeUsageEventHandler = async (event: CubeUsageEvent) => {
  const { SetUsage } = useRubiksCubeStore.getState();

  SetUsage(event.payload.cube_id, event.payload.usage);
};

export type IEvent =
  | CubeRotatedEvent
  | CubeAgentReasoningEvent
  | CubeAgentTimeoutEvent
  | CubeCompletedEvent
  | CubeUsageEvent;
