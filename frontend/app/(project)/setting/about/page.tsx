"use client";

import { useRouter } from "next/navigation";
import * as React from "react";

import { SettingsAbout } from "@/features/settings/components/sections/about/settings-about";
import { useFeaturePolicy } from "@/shared/hooks/use-feature-policy";

export default function SettingsAboutPage() {
  const router = useRouter();
  const { userAboutEnabled, loaded } = useFeaturePolicy();

  React.useEffect(() => {
    if (loaded && !userAboutEnabled) router.replace("/setting/general");
  }, [loaded, router, userAboutEnabled]);

  if (!loaded || !userAboutEnabled) return null;

  return <SettingsAbout />;
}
