import { getAIProviderIcon } from "@/components/(icons)/helpers";
import {
  Combobox,
  ComboboxContent,
  ComboboxEmpty,
  ComboboxInput,
  ComboboxItem,
  ComboboxList,
} from "@/components/ui/combobox";
import { InputGroupAddon } from "@/components/ui/input-group";
import type { LLM } from "@/domain/ai";

interface IProps {
  llms: LLM[];
  value?: LLM;
  onChange: (llm: LLM) => void;
}

export const LLMSelect = ({ llms, value, onChange }: IProps) => {
  const SelectedIcon = value ? getAIProviderIcon(value.provider) : null;

  return (
    <Combobox
      items={llms}
      value={value}
      itemToStringLabel={(llm: LLM) => llm.model}
      onValueChange={(value) => onChange(value as LLM)}
    >
      <ComboboxInput className="py-0" placeholder="Select model...">
        {SelectedIcon && (
          <InputGroupAddon align="inline-start">
            <SelectedIcon className="w-4 h-4" />
          </InputGroupAddon>
        )}
      </ComboboxInput>
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
