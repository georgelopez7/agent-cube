import { Link } from "@tanstack/react-router";

import { cn } from "cnfast";

import CubeLogo from "@/components/(icons)/cube-logo/cube-logo";
import GitHubIcon from "@/components/(icons)/github-icon/github-icon";
import XLogo from "@/components/(icons)/x-logo/x-logo";

interface IProps {
  className?: string;
  githubLink?: string;
  xLink?: string;
}

const Navbar = ({
  className,
  githubLink = "https://github.com",
  xLink = "https://x.com",
}: IProps) => {
  return (
    <nav className={cn("w-full", className)}>
      <div
        className={cn(
          "flex items-center justify-between max-w-6xl mx-auto px-6 py-2.5 border border-border bg-secondary/80 rounded-md",
        )}
      >
        <Link
          to="/"
          className="flex items-center gap-2 text-lg font-semibold tracking-tight text-foreground hover:text-primary transition-colors"
        >
          <CubeLogo size={28} />
          Agent Cube
        </Link>
        <div className="flex items-center gap-4">
          <a
            href={githubLink}
            target="_blank"
            rel="noopener noreferrer"
            className="text-white hover:text-white/80 transition-colors"
            aria-label="GitHub"
          >
            <GitHubIcon size={20} />
          </a>
          <a
            href={xLink}
            target="_blank"
            rel="noopener noreferrer"
            className="text-white hover:text-white/80 transition-colors"
            aria-label="X"
          >
            <XLogo size={18} />
          </a>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
