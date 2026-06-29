import { Plus } from "lucide-react";

import { Button } from "@/components/ui/button";
import { LLMSelect } from "@/components/(selects)/llm-select/llm-select";
import { AIProvider, type LLM } from "@/domain/ai";
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

const SAMPLE_LLMS: LLM[] = [
  { provider: AIProvider.OpenAI, model: "openai/gpt-4o" },
  { provider: AIProvider.Anthropic, model: "anthropic/claude-3-5-sonnet" },
  { provider: AIProvider.Google, model: "google/gemini-1.5-pro" },
];

const CreateAgentModal = () => {
  return (
    <Dialog>
      <DialogTrigger
        render={
          <Button className="w-fit cursor-pointer" variant="outline">
            <Plus />
            Create Agent
          </Button>
        }
      />
      <DialogContent className="border rounded-sm">
        <DialogHeader>
          <DialogTitle>Create Agent</DialogTitle>
          <DialogDescription>
            Select the model you want to use to create your agent.
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4">
          <div className="grid gap-1.5">
            <LLMSelect
              llms={SAMPLE_LLMS}
              onChange={(llm) => console.log(llm)}
            />
          </div>
        </div>
        <DialogFooter>
          <DialogClose render={<Button variant="outline">Cancel</Button>} />
          <Button>Create</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default CreateAgentModal;
