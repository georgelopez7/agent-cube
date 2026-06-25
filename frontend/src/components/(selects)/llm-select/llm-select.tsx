import {
  Combobox,
  ComboboxInput,
  ComboboxContent,
  ComboboxList,
  ComboboxItem,
  ComboboxEmpty,
} from "@/components/ui/combobox";
import type { LLM } from "@/domain/ai";
import { getAIProviderIcon } from "@/components/(icons)/helpers";

interface IProps {
  llms: LLM[];
  onChange: (llm: LLM) => void;
}

export const LLMSelect = ({ llms, onChange }: IProps) => {
  return (
    <Combobox
      items={llms}
      itemToStringLabel={(llm: LLM) => llm.model}
      onValueChange={(value) => onChange(value as LLM)}
    >
      <ComboboxInput className="py-0" placeholder="Select model..." />
      <ComboboxContent>
        <ComboboxEmpty>No models found.</ComboboxEmpty>
        <ComboboxList>
          {(item) => {
            const Icon = getAIProviderIcon(item.provider);
            return (
              <ComboboxItem key={item.model} value={item}>
                {Icon && <Icon className="w-4 h-4" />}
                {item.model}
              </ComboboxItem>
            );
          }}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  );
};
