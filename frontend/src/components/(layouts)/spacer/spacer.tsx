import { cn } from "cnfast";

interface IProps {
  className?: string;
  size?: "xxs" | "xs" | "sm" | "md" | "lg" | "xl";
}

const SIZE_CLASSES = {
  xxs: "py-1",
  xs: "py-2",
  sm: "py-4",
  md: "py-8",
  lg: "py-12",
  xl: "py-16",
};

const Spacer = ({ className = "", size = "md" }: IProps) => {
  return <div className={cn(SIZE_CLASSES[size], className)} />;
};

export default Spacer;
