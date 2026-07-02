import { useForm } from "@tanstack/react-form";
import { Plus } from "lucide-react";
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
import { AIProvider, type LLM } from "@/domain/ai";
import Spacer from "#/components/(layouts)/spacer/spacer";

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
});

export type CreateAgentFormData = z.infer<typeof schema>;

interface IProps {
  llms: LLM[];
  onSubmit: (data: CreateAgentFormData) => void | Promise<void>;
}

const CreateAgentModal = ({ llms, onSubmit }: IProps) => {
  const [open, setOpen] = useState(false);

  const form = useForm({
    defaultValues: {
      model: llms[0] ?? { provider: AIProvider.OpenAI, model: "" },
      scramble: 10,
    } satisfies CreateAgentFormData,
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
      <DialogTrigger
        render={
          <Button className="w-fit cursor-pointer" variant="outline">
            <Plus />
            Create Agent
          </Button>
        }
      />
      <DialogContent className="border rounded-sm">
        <form
          className="grid gap-6"
          onSubmit={(e) => {
            e.preventDefault();
            e.stopPropagation();
            form.handleSubmit();
          }}
        >
          <DialogHeader>
            <DialogTitle>Create Agent</DialogTitle>
            <DialogDescription>
              Select the model you want to use to create your agent.
            </DialogDescription>
          </DialogHeader>
          <div className="grid">
            <form.Field name="model">
              {(field) => (
                <div className="grid gap-1.5">
                  <label
                    htmlFor="model"
                    className="text-sm font-medium leading-none text-muted-foreground"
                  >
                    Model
                  </label>
                  <LLMSelect
                    llms={llms}
                    value={field.state.value}
                    onChange={(llm) => field.handleChange(llm)}
                  />
                  {field.state.meta.errors.length > 0 && (
                    <p className="text-sm text-destructive">
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
            <Spacer size="xs" />
            <Spacer size="xs" />
            <form.Field name="scramble">
              {(field) => (
                <div className="grid gap-1.5">
                  <label
                    htmlFor="scramble"
                    className="text-sm font-medium leading-none text-muted-foreground"
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
                    <p className="text-sm text-destructive">
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
            <Button type="submit" disabled={!form.state.canSubmit}>
              Create
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default CreateAgentModal;
