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
        "grid grid-cols-3 grid-rows-3 gap-0.5 p-0.5 bg-slate-900/80 rounded-md border border-slate-500/30",
        className,
      )}
      style={{ width: size, height: size }}
    >
      <div className="bg-rubiks-red rounded-[1px]" />
      <div className="bg-rubiks-blue rounded-[1px]" />
      <div className="bg-rubiks-green rounded-[1px]" />
      <div className="bg-rubiks-orange rounded-[1px]" />
      <div className="bg-rubiks-yellow rounded-[1px]" />
      <div className="bg-rubiks-white rounded-[1px]" />
      <div className="bg-rubiks-blue rounded-[1px]" />
      <div className="bg-rubiks-green rounded-[1px]" />
      <div className="bg-rubiks-yellow rounded-[1px]" />
    </div>
  );
};

export default CubeLogo;
