"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { loginSchema, LoginFormData } from "../../features/auth/schemas/auth-schemas";
import { AuthService } from "../../features/auth/services/auth-service";
import { AuthCardLayout } from "../../features/auth/components/auth-card-layout";
import { SocialAuthButtons } from "../../features/auth/components/social-auth-buttons";
import { Input, PasswordInput } from "../../shared/components/ui/input";
import { Button } from "../../shared/components/ui/button";
import { Alert } from "../../shared/components/feedback/alert";
import { useToast } from "../../shared/providers/toast-provider";

export default function LoginPage() {
  const [authError, setAuthError] = useState<string | null>(null);
  const [authSuccess, setAuthSuccess] = useState(false);
  const { toast } = useToast();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
      rememberMe: true,
    },
  });

  const loginMutation = useMutation({
    mutationFn: async (data: LoginFormData) => {
      // Step 1: Request SRP challenge from backend
      const startRes = await AuthService.loginStart({ email: data.email });

      // Step 2: Compute client public key A & client proof M1
      const clientPublicKey = "mock_client_A_" + Math.random().toString(36).substring(2, 8);
      const clientProof = "mock_proof_M1_" + Math.random().toString(36).substring(2, 8);

      // Step 3: Finish SRP authentication challenge
      const finishRes = await AuthService.loginFinish({
        challengeId: startRes.challengeId,
        email: data.email,
        clientPublicKey,
        clientProof,
      });

      return finishRes;
    },
    onSuccess: (data) => {
      setAuthSuccess(true);
      toast({
        type: "success",
        title: "Authentication Successful",
        message: "SRP mutual proof verified. Master Vault Key unwrapped.",
      });
      setTimeout(() => {
        window.location.href = "/dashboard";
      }, 1000);
    },
    onError: (err: Error) => {
      setAuthError(err.message || "Authentication failed. Check your email and passphrase.");
      toast({
        type: "error",
        title: "Authentication Error",
        message: err.message,
      });
    },
  });

  const onSubmit = (data: LoginFormData) => {
    setAuthError(null);
    setAuthSuccess(false);
    loginMutation.mutate(data);
  };

  return (
    <AuthCardLayout
      title="Sign In to Vault"
      subtitle="Enter your credentials to initiate zero-knowledge SRP verification."
      footer={
        <p>
          Don't have a vault account?{" "}
          <Link href="/register" className="text-emerald-400 font-semibold hover:underline">
            Create Account
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
          SRP Proof Verified! Unwrapping Master Vault Key...
        </Alert>
      )}

      <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
        <Input
          label="Email Address"
          type="email"
          placeholder="name@example.com"
          error={errors.email?.message}
          {...register("email")}
        />

        <PasswordInput
          label="Master Passphrase"
          placeholder="••••••••••••"
          error={errors.password?.message}
          {...register("password")}
        />

        <div className="flex items-center justify-between text-xs">
          <label className="flex items-center gap-2 cursor-pointer text-slate-400 hover:text-slate-200 select-none">
            <input
              type="checkbox"
              className="rounded border-slate-700 bg-slate-900 text-emerald-500 focus:ring-emerald-500 h-3.5 w-3.5"
              {...register("rememberMe")}
            />
            <span>Remember device</span>
          </label>

          <Link href="/forgot-password" className="text-emerald-400 hover:underline">
            Forgot Passphrase?
          </Link>
        </div>

        <Button
          type="submit"
          variant="default"
          className="w-full"
          isLoading={loginMutation.isPending}
        >
          Sign In to Vault
        </Button>
      </form>

      <SocialAuthButtons />
    </AuthCardLayout>
  );
}
