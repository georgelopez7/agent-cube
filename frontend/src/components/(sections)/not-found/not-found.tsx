import { Link } from "@tanstack/react-router";
import { Home } from "lucide-react";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { buttonVariants } from "#/components/ui/button";

const NotFound = () => {
  return (
    <PageLayout>
      <div className="flex flex-1 flex-col items-center justify-center text-center">
        <h1 className="text-8xl md:text-9xl font-bold font-heading tracking-tight text-foreground">
          404
        </h1>
        <Spacer size="sm" />
        <p className="text-lg md:text-xl text-muted-foreground max-w-md leading-relaxed">
          You&apos;ve gotten all mixed up... let&apos;s get you back to being
          solved.
        </p>
        <Spacer size="sm" />
        <Link to="/" className={buttonVariants()}>
          <Home />
          Back to home
        </Link>
      </div>
    </PageLayout>
  );
};

export default NotFound;
