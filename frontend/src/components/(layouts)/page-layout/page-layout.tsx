import Navbar from "#/components/(layouts)/navbar/navbar";

interface IPageLayoutProps {
  children: React.ReactNode;
  navbar?: boolean;
}

export const PageLayout = ({ children, navbar = true }: IPageLayoutProps) => {
  return (
    <div className="relative flex flex-col min-h-screen p-8 mx-auto">
      {navbar && <Navbar />}
      {/* Top Left Corner - Green */}
      <div className="absolute top-4 left-4 w-8 h-8">
        <div className="absolute top-0 left-0 w-full h-0.5 bg-rubiks-green" />
        <div className="absolute top-0 left-0 w-0.5 h-full bg-rubiks-green" />
      </div>

      {/* Top Right Corner - Blue */}
      <div className="absolute top-4 right-4 w-8 h-8">
        <div className="absolute top-0 right-0 w-full h-0.5 bg-rubiks-blue" />
        <div className="absolute top-0 right-0 w-0.5 h-full bg-rubiks-blue" />
      </div>

      {/* Bottom Left Corner - Yellow */}
      <div className="absolute bottom-4 left-4 w-8 h-8">
        <div className="absolute bottom-0 left-0 w-full h-0.5 bg-rubiks-yellow" />
        <div className="absolute bottom-0 left-0 w-0.5 h-full bg-rubiks-yellow" />
      </div>

      {/* Bottom Right Corner - Red */}
      <div className="absolute bottom-4 right-4 w-8 h-8">
        <div className="absolute bottom-0 right-0 w-full h-0.5 bg-rubiks-red" />
        <div className="absolute bottom-0 right-0 w-0.5 h-full bg-rubiks-red" />
      </div>

      {children}
    </div>
  );
};
