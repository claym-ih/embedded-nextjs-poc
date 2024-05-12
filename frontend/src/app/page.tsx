"use client";

import { Suspense } from "react";

import { WelcomePage } from "@refinedev/core";
import UserPage from "./user/page";
import {NavigateToResource} from "@refinedev/nextjs-router";

export default function IndexPage() {
  return (
    <Suspense>
        <NavigateToResource />
    </Suspense>
  );
}
