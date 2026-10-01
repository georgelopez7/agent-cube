import cn from "cnfast";
import Navbar from "#/components/(layouts)/navbar/navbar";

interface IPageLayoutProps {
  children: React.ReactNode;
  navbar?: boolean;
  showGrid?: boolean;
  background?: React.ReactNode;
}

export const PageLayout = ({
  children,
  navbar = true,
  showGrid = true,
  background,
}: IPageLayoutProps) => {
  return (
    <div
      className={cn(
        "relative flex min-h-screen flex-col overflow-hidden bg-black p-4 text-white sm:p-6",
        showGrid && "blueprint-grid",
      )}
    >
      {background}
      <div className="relative z-10 mx-auto w-full max-w-6xl">
        {navbar && (
          <Navbar
            githubLink={import.meta.env.VITE_GITHUB_URL}
            xLink={import.meta.env.VITE_X_URL}
          />
        )}
      </div>
      {/* Top Left Corner - Green */}
      <div className="absolute top-4 left-4 w-8 h-8 pointer-events-none">
        <div className="absolute top-0 left-0 w-full h-0.5 bg-rubiks-green" />
        <div className="absolute top-0 left-0 w-0.5 h-full bg-rubiks-green" />
      </div>

      {/* Top Right Corner - Blue */}
      <div className="absolute top-4 right-4 w-8 h-8 pointer-events-none">
        <div className="absolute top-0 right-0 w-full h-0.5 bg-rubiks-blue" />
        <div className="absolute top-0 right-0 w-0.5 h-full bg-rubiks-blue" />
      </div>

      {/* Bottom Left Corner - Yellow */}
      <div className="absolute bottom-4 left-4 w-8 h-8 pointer-events-none">
        <div className="absolute bottom-0 left-0 w-full h-0.5 bg-rubiks-yellow" />
        <div className="absolute bottom-0 left-0 w-0.5 h-full bg-rubiks-yellow" />
      </div>

      {/* Bottom Right Corner - Red */}
      <div className="absolute bottom-4 right-4 w-8 h-8 pointer-events-none">
        <div className="absolute bottom-0 right-0 w-full h-0.5 bg-rubiks-red" />
        <div className="absolute bottom-0 right-0 w-0.5 h-full bg-rubiks-red" />
      </div>

      {/* Top Center - White */}
      <div className="absolute top-4 left-1/2 -translate-x-1/2 w-8 h-0.5 bg-rubiks-white pointer-events-none" />

      {/* Bottom Center - Orange */}
      <div className="absolute bottom-4 left-1/2 -translate-x-1/2 w-8 h-0.5 bg-rubiks-orange pointer-events-none" />

      <main className="relative z-10 mx-auto flex w-full max-w-6xl flex-1 flex-col">
        {children}
      </main>
    </div>
  );
};
