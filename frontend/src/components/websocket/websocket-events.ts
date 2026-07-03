import { RubiksCubeStatus, type CubeRotation } from "#/domain/rubiks-cube";
import { useRubiksCubeStore } from "#/stores/rubiks-cube-store";

export enum WebSocketEventType {
  CubeRotated = "cube.rotated",
  CubeAgentTimeout = "cube.agent.timeout",
  CubeCompleted = "cube.completed",
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

export type IEvent =
  | CubeRotatedEvent
  | CubeAgentTimeoutEvent
  | CubeCompletedEvent;
