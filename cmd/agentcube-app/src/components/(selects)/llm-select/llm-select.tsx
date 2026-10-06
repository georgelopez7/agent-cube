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
  const selectedValue = value?.model ? value : null;
  const SelectedIcon = selectedValue
    ? getAIProviderIcon(selectedValue.provider)
    : null;

  return (
    <Combobox
      items={llms}
      value={selectedValue}
      itemToStringLabel={(llm: LLM) => llm.model}
      onValueChange={(value) => {
        if (value) {
          onChange(value as LLM);
        }
      }}
    >
      <ComboboxInput
        id="model"
        inputClassName="py-0 text-xs"
        placeholder="Select model"
      >
        {SelectedIcon && (
          <InputGroupAddon align="inline-start" className="h-full py-0 pl-3">
            <SelectedIcon className="size-4 shrink-0" />
          </InputGroupAddon>
        )}
      </ComboboxInput>
      <ComboboxContent align="start" sideOffset={4}>
        <ComboboxEmpty>No models found.</ComboboxEmpty>
        <ComboboxList>
          {(item) => {
            const Icon = getAIProviderIcon(item.provider);
            return (
              <ComboboxItem key={item.model} value={item}>
                {Icon && <Icon className="size-4 shrink-0" />}
                <span className="truncate">{item.model}</span>
              </ComboboxItem>
            );
          }}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  );
};
