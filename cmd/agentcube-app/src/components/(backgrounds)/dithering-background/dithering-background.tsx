import { Dithering } from "@paper-design/shaders-react";
import cn from "cnfast";

interface IProps {
  className?: string;
}

const DitheringBackground = ({ className }: IProps) => {
  return (
    <div
      aria-hidden="true"
      className={cn(
        "pointer-events-none absolute inset-0 overflow-hidden",
        className,
      )}
    >
      <Dithering
        width="100%"
        height="100%"
        colorBack="#101010"
        colorFront="#737373"
        shape="warp"
        type="4x4"
        size={2.5}
        speed={0.8}
        scale={0.8}
        fit="cover"
      />
      <div className="absolute inset-0 bg-black/60" />
    </div>
  );
};

export default DitheringBackground;
