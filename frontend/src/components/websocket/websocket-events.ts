import type { CubeRotation } from "#/domain/rubiks-cube";
import { useRubiksCubeStore } from "#/stores/rubiks-cube-store";

export enum WebSocketEventType {
  CubeRotated = "cube.rotated",
  // CubeSolved = "cube.solved",
  // AgentMessage = "agent.message",
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

// type CubeSolvedEvent = {
//   type: "cube.solved";
//   payload: {
//     cube_id: string;
//   };
// };

// export const SolvedEventHandler = async (event: CubeSolvedEvent) => {
//   const { SetStatus } = useRubiksCubeStore.getState();

//   SetStatus(event.payload.cube_id, IStatus.Completed);
// };

// type AgentMessageEvent = {
//   type: "agent.message";
//   payload: {
//     cube_id: string;
//     message: IAgentMessage;
//   };
// };

// export const AgentMessageEventHandler = async (event: AgentMessageEvent) => {
//   const { AddAgentMessage } = useRubiksCubeStore.getState();

//   AddAgentMessage(event.payload.cube_id, event.payload.message);
// };

export type IEvent = CubeRotatedEvent;
