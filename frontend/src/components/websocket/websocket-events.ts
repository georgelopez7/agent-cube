export type RotationEvent = {
  event_type: "cube.rotation";
  id: string;
};

export const RotationEventHandler = async (event: RotationEvent) => {
  console.log(event);
  // const { records, AddRotation } = useRubiksCubeStore.getState();

  // const record = records.find((record) => record.record?._id === event.id);
  // if (!record || !record.cube) return;

  // await record.cube.rotate(event.rotation.rotation); // ROTATE CUBE

  // AddRotation(event.id, event.rotation);
};

// type ISolvedEvent = {
//   event_type: "cube.solved";
//   id: string;
// };

// export const SolvedEventHandler = async (event: ISolvedEvent) => {
//   const { SetStatus } = useRubiksCubeStore.getState();

//   SetStatus(event.id, IStatus.Completed);
// };

// type IAgentMessageEvent = {
//   event_type: "agent.message";
//   id: string;
//   message: IAgentMessage;
// };

// export const AgentMessageEventHandler = async (event: IAgentMessageEvent) => {
//   const { AddAgentMessage } = useRubiksCubeStore.getState();

//   AddAgentMessage(event.id, event.message);
// };

export type IEvent = RotationEvent;
