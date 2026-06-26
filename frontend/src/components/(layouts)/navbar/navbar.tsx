import { Link } from "@tanstack/react-router";

import { cn } from "cnfast";

import { buttonVariants } from "@/components/ui/button";

interface IProps {
  className?: string;
}

const Navbar = ({ className }: IProps) => {
  return (
    <nav
      className={cn(
        "flex items-center justify-between w-full px-6 py-4 border border-border bg-background",
        className,
      )}
    >
      <Link
        to="/"
        className="text-lg font-semibold tracking-tight text-foreground hover:text-primary transition-colors"
      >
        Agent Cube
      </Link>
      <a
        href="https://github.com"
        target="_blank"
        rel="noopener noreferrer"
        className={buttonVariants()}
      >
        GitHub
      </a>
    </nav>
  );
};

export default Navbar;
