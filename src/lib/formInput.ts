import { ChangeEvent, useState } from "react";


interface FormState {
    [key: string]: string;
}
  
// Utility function for handling form input changes
export const useFormInput = <T extends FormState>(
    initialState: T
    ): [T, (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void] => {
    const [state, setState] = useState<T>(initialState);
  
    const handleForm = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>): void => {
      setState((prev) => ({
        ...prev,
        [e.target.name]: e.target.value,
      }));
    };
  
    return [state, handleForm];
};

