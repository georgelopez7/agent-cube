import { useForm } from "@tanstack/react-form";
import { AlertTriangle, Plus } from "lucide-react";
import { useState } from "react";
import { z } from "zod";
import { LLMSelect } from "@/components/(selects)/llm-select/llm-select";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { AIProvider, type LLM } from "@/domain/ai";

const MAX_DURATION_OPTIONS = [
  { label: "15 seconds", value: 15000 },
  { label: "30 seconds", value: 30000 },
  { label: "1 minute", value: 60000 },
  { label: "5 minutes", value: 300000 },
  { label: "10 minutes", value: 600000 },
  { label: "Infinity", value: 86400000 },
];

const schema = z.object({
  model: z.object({
    provider: z.enum(AIProvider),
    model: z.string().min(1, "Model is required"),
  }),
  scramble: z
    .number({ message: "Scramble must be a number" })
    .int("Scramble must be a whole number")
    .min(0, "Scramble must be at least 0")
    .max(100, "Scramble must be at most 100"),
  maxDuration: z.number({ message: "Max duration must be a number" }),
});

export type NewAgentFormData = z.infer<typeof schema>;

interface IProps {
  llms: LLM[];
  onSubmit: (data: NewAgentFormData) => void | Promise<void>;
}

const NewAgentDialog = ({ llms, onSubmit }: IProps) => {
  const [open, setOpen] = useState(false);

  const form = useForm({
    defaultValues: {
      model: { provider: AIProvider.OpenAI, model: "" },
      scramble: 10,
      maxDuration: 60000,
    } as NewAgentFormData,
    validators: {
      onChange: schema,
    },
    onSubmit: async ({ value }) => {
      await onSubmit(value);
      setOpen(false);
    },
  });

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <div className="w-fit bg-[#101010] p-2">
        <DialogTrigger
          render={
            <Button className="w-fit cursor-pointer" variant="outline">
              <Plus />
              New Agent
            </Button>
          }
        />
      </div>
      <DialogContent
        className="bg-[#101010] p-3 ring-0 sm:max-w-md"
        showCloseButton={false}
      >
        <form
          className="grid gap-6"
          onSubmit={(e) => {
            e.preventDefault();
            e.stopPropagation();
            form.handleSubmit();
          }}
        >
          <DialogHeader className="border border-neutral-700 bg-black px-4 py-3">
            <DialogTitle>New Agent</DialogTitle>
            <DialogDescription>
              Select the model you want to use to create your agent.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4">
            <form.Field name="model">
              {(field) => (
                <div className="grid gap-1.5">
                  <label
                    htmlFor="model"
                    className="font-mono text-xs uppercase tracking-wide text-neutral-400"
                  >
                    Model
                  </label>
                  <LLMSelect
                    llms={llms}
                    value={field.state.value}
                    onChange={(llm) => field.handleChange(llm)}
                  />
                  {field.state.meta.errors.length > 0 && (
                    <p className="font-mono text-xs text-red-500">
                      {field.state.meta.errors
                        .map((error) =>
                          typeof error === "string" ? error : error?.message,
                        )
                        .filter(Boolean)
                        .join(", ")}
                    </p>
                  )}
                </div>
              )}
            </form.Field>
            <form.Field name="scramble">
              {(field) => (
                <div className="grid gap-1.5">
                  <label
                    htmlFor="scramble"
                    className="font-mono text-xs uppercase tracking-wide text-neutral-400"
                  >
                    Scramble
                  </label>
                  <Input
                    id="scramble"
                    type="number"
                    value={field.state.value}
                    onChange={(e) => field.handleChange(e.target.valueAsNumber)}
                  />
                  {field.state.meta.errors.length > 0 && (
                    <p className="font-mono text-xs text-red-500">
                      {field.state.meta.errors
                        .map((error) =>
                          typeof error === "string" ? error : error?.message,
                        )
                        .filter(Boolean)
                        .join(", ")}
                    </p>
                  )}
                </div>
              )}
            </form.Field>
            <form.Field name="maxDuration">
              {(field) => (
                <div className="grid gap-1.5">
                  <label
                    htmlFor="maxDuration"
                    className="font-mono text-xs uppercase tracking-wide text-neutral-400"
                  >
                    Max Duration
                  </label>
                  <Select
                    value={String(field.state.value)}
                    onValueChange={(value) => field.handleChange(Number(value))}
                  >
                    <SelectTrigger id="maxDuration" className="w-full">
                      <SelectValue placeholder="Select max duration">
                        {MAX_DURATION_OPTIONS.find(
                          (option) => option.value === field.state.value,
                        )?.label ?? "Select max duration"}
                      </SelectValue>
                    </SelectTrigger>
                    <SelectContent>
                      {MAX_DURATION_OPTIONS.map((option) => (
                        <SelectItem
                          key={option.value}
                          value={String(option.value)}
                        >
                          {option.label}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  {field.state.value === 86400000 && (
                    <p className="inline-flex items-center gap-1.5 font-mono text-xs text-yellow-500">
                      <AlertTriangle className="size-4" />
                      Infinity means this agent could continue forever.
                    </p>
                  )}
                  {field.state.meta.errors.length > 0 && (
                    <p className="font-mono text-xs text-red-500">
                      {field.state.meta.errors
                        .map((error) =>
                          typeof error === "string" ? error : error?.message,
                        )
                        .filter(Boolean)
                        .join(", ")}
                    </p>
                  )}
                </div>
              )}
            </form.Field>
          </div>
          <DialogFooter>
            <DialogClose
              render={
                <Button type="button" variant="outline">
                  Cancel
                </Button>
              }
            />
            <Button
              type="submit"
              className="border-neutral-700 bg-white text-black hover:bg-neutral-200"
              disabled={!form.state.canSubmit}
            >
              Create
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default NewAgentDialog;
