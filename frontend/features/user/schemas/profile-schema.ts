import { z } from "zod";

export const profileSchema = z.object({
  displayName: z.string().min(2, "Display name must be at least 2 characters").max(50),
  avatarUrl: z.string().optional(),
  timezone: z.string().min(1, "Timezone is required"),
  language: z.string().min(1, "Language is required"),
  bio: z.string().max(300, "Bio cannot exceed 300 characters").optional(),
});

export type ProfileFormData = z.infer<typeof profileSchema>;

export const preferencesSchema = z.object({
  theme: z.enum(["dark", "light", "system"]),
  aiBehavior: z.enum(["minimal", "balanced", "proactive"]),
  privacyMode: z.enum(["local_first", "encrypted_cloud", "hybrid"]),
  emailNotifications: z.boolean(),
  pushNotifications: z.boolean(),
  desktopNotifications: z.boolean(),
  securityAlerts: z.boolean(),
});

export type PreferencesFormData = z.infer<typeof preferencesSchema>;
