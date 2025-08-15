import { ReactNode } from 'react';


export type ButtonVariant = 'primary' | 'secondary' | 'ghost' | 'destructive';
export type ButtonSize = 'sm' | 'md' | 'lg';

export interface ButtonProps {
  variant?: ButtonVariant;
  size?: ButtonSize;
  disabled?: boolean;
  className?: string;
  children?: ReactNode;
  onClick?: () => void;
}

export interface InputProps {
  id?: string;
  name?: string;
  value?: string;
  placeholder?: string;
  disabled?: boolean;
  className?: string;
  onChange?: (value: string) => void;
}

export interface ModalProps {
  isOpen: boolean;
  title?: string;
  onClose: () => void;
  children?: ReactNode;
}