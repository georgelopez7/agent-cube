import { createFileRoute } from "@tanstack/react-router";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import CreateAgentModal from "#/components/(modals)/create-agent-modal/create-agent-modal";

export const Route = createFileRoute("/(pages)/agents")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <PageLayout>
      <div className="w-[80vw] flex items-center justify-between gap-4">
        <h1 className="text-4xl font-semibold">Agents</h1>
        <CreateAgentModal />
      </div>
    </PageLayout>
  );
}
