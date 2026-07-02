import { createFileRoute } from "@tanstack/react-router";

import NotFound from "#/components/(sections)/not-found/not-found";

export const Route = createFileRoute("/$")({
  component: NotFound,
});
