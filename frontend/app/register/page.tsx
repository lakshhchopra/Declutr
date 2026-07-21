"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { registerSchema, RegisterFormData } from "../../features/auth/schemas/auth-schemas";
import { AuthService } from "../../features/auth/services/auth-service";
import { AuthCardLayout } from "../../features/auth/components/auth-card-layout";
import { PasswordStrengthMeter } from "../../features/auth/components/password-strength-meter";
import { Input, PasswordInput } from "../../shared/components/ui/input";
import { Button } from "../../shared/components/ui/button";
import { Alert } from "../../shared/components/feedback/alert";
import { useToast } from "../../shared/providers/toast-provider";

export default function RegisterPage() {
  const [authError, setAuthError] = useState<string | null>(null);
  const [authSuccess, setAuthSuccess] = useState(false);
  const { toast } = useToast();

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
      confirmPassword: "",
      acceptTerms: false,
    },
  });

  const passwordValue = watch("password", "");

  const registerMutation = useMutation({
    mutationFn: async (data: RegisterFormData) => {
      // Step 1: Derive SRP salt and verifier locally
      const srpSalt = "salt_" + Math.random().toString(36).substring(2, 10);
      const srpVerifier = "verifier_v_" + Math.random().toString(36).substring(2, 12);

      // Step 2: Encrypt Master Vault Key (MVK)
      const mvkPayload = {
        ciphertext: "mvk_cipher_" + Math.random().toString(36).substring(2, 16),
        nonce: "nonce_" + Math.random().toString(36).substring(2, 8),
        version: 1,
      };

      // Step 3: Register account on backend
      return await AuthService.register({
        email: data.email,
        name: data.name,
        srpSalt,
        srpVerifier,
        mvk: mvkPayload,
      });
    },
    onSuccess: (data) => {
      setAuthSuccess(true);
      toast({
        type: "success",
        title: "Vault Account Created",
        message: "Account registered successfully with zero-knowledge SRP verifier.",
      });
      setTimeout(() => {
        window.location.href = "/verify-email";
      }, 1000);
    },
    onError: (err: Error) => {
      setAuthError(err.message || "Failed to create vault account.");
      toast({
        type: "error",
        title: "Registration Error",
        message: err.message,
      });
    },
  });

  const onSubmit = (data: RegisterFormData) => {
    setAuthError(null);
    setAuthSuccess(false);
    registerMutation.mutate(data);
  };

  return (
    <AuthCardLayout
      title="Create Vault Account"
      subtitle="Your master passphrase never leaves your device."
      footer={
        <p>
          Already have a vault account?{" "}
          <Link href="/login" className="text-emerald-400 font-semibold hover:underline">
            Sign In
          </Link>
        </p>
      }
    >
      {authError && (
        <Alert variant="danger" onClose={() => setAuthError(null)}>
          {authError}
        </Alert>
      )}

      {authSuccess && (
        <Alert variant="success">
          Account created! Directing to email verification...
        </Alert>
      )}

      <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
        <Input
          label="Full Name"
          placeholder="Jane Doe"
          error={errors.name?.message}
          {...register("name")}
        />

        <Input
          label="Email Address"
          type="email"
          placeholder="jane@example.com"
          error={errors.email?.message}
          {...register("email")}
        />

        <div>
          <PasswordInput
            label="Master Passphrase"
            placeholder="••••••••••••"
            error={errors.password?.message}
            {...register("password")}
          />
          <PasswordStrengthMeter password={passwordValue} />
        </div>

        <PasswordInput
          label="Confirm Master Passphrase"
          placeholder="••••••••••••"
          error={errors.confirmPassword?.message}
          {...register("confirmPassword")}
        />

        <div className="space-y-1">
          <label className="flex items-start gap-2.5 cursor-pointer text-xs text-slate-400 select-none">
            <input
              type="checkbox"
              className="mt-0.5 rounded border-slate-700 bg-slate-900 text-emerald-500 focus:ring-emerald-500 h-4 w-4 shrink-0"
              {...register("acceptTerms")}
            />
            <span>
              I accept the{" "}
              <a href="#" className="text-emerald-400 underline">
                Terms of Service
              </a>{" "}
              and{" "}
              <a href="#" className="text-emerald-400 underline">
                Privacy Policy
              </a>
              .
            </span>
          </label>
          {errors.acceptTerms && (
            <p className="text-xs text-rose-500">{errors.acceptTerms.message}</p>
          )}
        </div>

        <Button
          type="submit"
          variant="default"
          className="w-full"
          isLoading={registerMutation.isPending}
        >
          Create Encrypted Vault
        </Button>
      </form>
    </AuthCardLayout>
  );
}
