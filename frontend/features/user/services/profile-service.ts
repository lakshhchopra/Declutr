import { ProfileFormData, PreferencesFormData } from "../schemas/profile-schema";

export interface UserProfile extends ProfileFormData {
  onboardingCompleted: boolean;
}

export interface UserPreferences extends PreferencesFormData {}

const PROFILE_STORAGE_KEY = "declutr_user_profile";
const PREFS_STORAGE_KEY = "declutr_user_preferences";

export const ProfileService = {
  async getProfile(): Promise<UserProfile> {
    try {
      const stored = localStorage.getItem(PROFILE_STORAGE_KEY);
      if (stored) return JSON.parse(stored);
    } catch (err) {}
    return {
      displayName: "Declutr User",
      timezone: "UTC",
      language: "en",
      onboardingCompleted: false,
    };
  },

  async updateProfile(data: Partial<UserProfile>): Promise<UserProfile> {
    const current = await this.getProfile();
    const updated = { ...current, ...data };
    localStorage.setItem(PROFILE_STORAGE_KEY, JSON.stringify(updated));
    return updated;
  },

  async getPreferences(): Promise<UserPreferences> {
    try {
      const stored = localStorage.getItem(PREFS_STORAGE_KEY);
      if (stored) return JSON.parse(stored);
    } catch (err) {}
    return {
      theme: "dark",
      aiBehavior: "balanced",
      privacyMode: "local_first",
      emailNotifications: true,
      pushNotifications: true,
      desktopNotifications: true,
      securityAlerts: true,
    };
  },

  async updatePreferences(data: Partial<UserPreferences>): Promise<UserPreferences> {
    const current = await this.getPreferences();
    const updated = { ...current, ...data };
    localStorage.setItem(PREFS_STORAGE_KEY, JSON.stringify(updated));
    return updated;
  },

  async completeOnboarding(): Promise<void> {
    await this.updateProfile({ onboardingCompleted: true });
  },
};
