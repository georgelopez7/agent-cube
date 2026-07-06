import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createRootRoute, HeadContent, Scripts } from "@tanstack/react-router";

// import { TanStackRouterDevtoolsPanel } from "@tanstack/react-router-devtools";
// import { TanStackDevtools } from "@tanstack/react-devtools";

import { Toaster } from "sonner";

import "#/styles.css";
import { getOGImage, getSiteUrl } from "#/utils/seo";

const queryClient = new QueryClient();

const SEO = {
  url: getSiteUrl(),
  title: "Agent Cube",
  description:
    "Benchmark large language models by challenging them to solve a Rubik's Cube.",
  image: getOGImage(),
};

export const Route = createRootRoute({
  head: () => ({
    meta: [
      {
        title: SEO.title,
      },
      {
        name: "description",
        content: SEO.description,
      },
      {
        property: "og:title",
        content: SEO.title,
      },
      {
        property: "og:description",
        content: SEO.description,
      },
      {
        property: "og:type",
        content: "website",
      },
      {
        property: "og:image",
        content: SEO.image,
      },
      {
        name: "twitter:card",
        content: "summary_large_image",
      },
      {
        name: "twitter:title",
        content: SEO.title,
      },
      {
        name: "twitter:description",
        content: SEO.description,
      },
      {
        name: "twitter:image",
        content: SEO.image,
      },
    ],
  }),
  shellComponent: RootDocument,
});

function RootDocument({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>
        <HeadContent />
      </head>
      <body>
        <QueryClientProvider client={queryClient}>
          {children}
          <Toaster richColors />
        </QueryClientProvider>
        {/* <TanStackDevtools
          config={{
            position: 'bottom-right',
          }}
          plugins={[
            {
              name: 'Tanstack Router',
              render: <TanStackRouterDevtoolsPanel />,
            },
          ]}
        /> */}
        <Scripts />
      </body>
    </html>
  );
}
