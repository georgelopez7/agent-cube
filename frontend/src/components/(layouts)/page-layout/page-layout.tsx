import Navbar from "#/components/(layouts)/navbar/navbar";
import Spacer from "../spacer/spacer";

interface IPageLayoutProps {
  children: React.ReactNode;
  navbar?: boolean;
}

export const PageLayout = ({ children, navbar = true }: IPageLayoutProps) => {
  return (
    <div className="relative flex flex-col min-h-screen p-8 mx-auto">
      <Spacer size="sm" />
      <div className="relative z-10">
        {navbar && (
          <Navbar
            githubLink={import.meta.env.VITE_GITHUB_URL}
            xLink={import.meta.env.VITE_X_URL}
          />
        )}
      </div>
      <Spacer size="xs" />
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

      <main className="flex flex-1 flex-col w-full max-w-6xl mx-auto">
        {children}
      </main>
    </div>
  );
};
