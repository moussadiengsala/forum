import { ChangeEvent, useState } from "react";

interface FormState {
    [key: string]: string | File | null;
}

// Utility function for handling form input changes
export const useFormInput = <T extends FormState>(
    initialState: T
  ): [T, (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void] => {
  const [state, setState] = useState<T>(initialState);

  const handleForm = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>): void => {
    const { name, type, value } = e.target;

    setState((prev) => ({
      ...prev,
      [name]: type === "file" ? (e.target as HTMLInputElement).files?.[0] : value,
    }));
  };

  return [state, handleForm];
};
