import { cn } from "cnfast";

interface IProps {
  className?: string;
  size?: number;
}

const CubeLogo = ({ className, size = 28 }: IProps) => {
  return (
    <div
      role="img"
      aria-label="Agent Cube logo"
      className={cn(
        "grid grid-cols-3 grid-rows-3 gap-0.5 bg-[#101010] p-0.5",
        className,
      )}
      style={{ width: size, height: size }}
    >
      <div className="bg-rubiks-red" />
      <div className="bg-rubiks-blue" />
      <div className="bg-rubiks-green" />
      <div className="bg-rubiks-orange" />
      <div className="bg-rubiks-yellow" />
      <div className="bg-rubiks-white" />
      <div className="bg-rubiks-blue" />
      <div className="bg-rubiks-green" />
      <div className="bg-rubiks-yellow" />
    </div>
  );
};

export default CubeLogo;
