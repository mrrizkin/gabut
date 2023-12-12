import { z } from "zod";

export const authFormSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

export type AuthForm = z.infer<typeof authFormSchema>;
