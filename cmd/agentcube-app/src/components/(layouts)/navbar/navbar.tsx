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
    <nav className={cn("mt-4 mb-2 w-full sm:mt-6 sm:mb-4", className)}>
      <div className="bg-[#101010] p-3">
        <div className="flex items-center justify-between border border-neutral-700 bg-black px-4 py-3 font-mono">
          <Link
            to="/"
            className="group flex items-center gap-2 font-heading text-xl font-extrabold uppercase tracking-tight text-white hover:underline hover:decoration-[3px] hover:underline-offset-4"
          >
            <CubeLogo
              size={24}
              className="group-hover:animate-[spin_0.25s_linear_8]"
            />
            Agent Cube
          </Link>
          <div className="flex items-center gap-2">
            <a
              href={githubLink}
              target="_blank"
              rel="noopener noreferrer"
              className="flex size-9 cursor-pointer items-center justify-center border border-transparent text-white transition-colors hover:border-neutral-700 hover:bg-neutral-900 hover:text-white"
              aria-label="GitHub"
            >
              <GitHubIcon size={20} />
            </a>
            <a
              href={xLink}
              target="_blank"
              rel="noopener noreferrer"
              className="flex size-9 cursor-pointer items-center justify-center border border-transparent text-white transition-colors hover:border-neutral-700 hover:bg-neutral-900 hover:text-white"
              aria-label="X"
            >
              <XLogo size={18} />
            </a>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
