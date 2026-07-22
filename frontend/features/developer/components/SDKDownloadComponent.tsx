"use client";

import React from "react";
import { Download, Terminal, Code, Package } from "lucide-react";

export function SDKDownloadComponent() {
  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-xl font-bold tracking-tight">Official SDKs & Declutr CLI</h3>
        <p className="text-sm text-muted-foreground">Download client libraries and command line tool binaries</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <div className="p-5 rounded-xl border bg-card space-y-3">
          <div className="p-2.5 rounded-lg bg-indigo-500/10 text-indigo-500 w-fit">
            <Code className="w-5 h-5" />
          </div>
          <h4 className="font-bold text-sm">TypeScript / Node.js SDK</h4>
          <p className="text-xs text-muted-foreground font-mono">npm install @declutr/sdk</p>
          <button className="w-full py-1.5 rounded-lg border text-xs font-semibold hover:bg-secondary transition-colors">
            View Package
          </button>
        </div>

        <div className="p-5 rounded-xl border bg-card space-y-3">
          <div className="p-2.5 rounded-lg bg-emerald-500/10 text-emerald-500 w-fit">
            <Package className="w-5 h-5" />
          </div>
          <h4 className="font-bold text-sm">Go Client Library</h4>
          <p className="text-xs text-muted-foreground font-mono">go get github.com/declutr/sdks/go</p>
          <button className="w-full py-1.5 rounded-lg border text-xs font-semibold hover:bg-secondary transition-colors">
            View Package
          </button>
        </div>

        <div className="p-5 rounded-xl border bg-card space-y-3">
          <div className="p-2.5 rounded-lg bg-amber-500/10 text-amber-500 w-fit">
            <Code className="w-5 h-5" />
          </div>
          <h4 className="font-bold text-sm">Python Package</h4>
          <p className="text-xs text-muted-foreground font-mono">pip install declutr-sdk</p>
          <button className="w-full py-1.5 rounded-lg border text-xs font-semibold hover:bg-secondary transition-colors">
            View Package
          </button>
        </div>

        <div className="p-5 rounded-xl border bg-card space-y-3">
          <div className="p-2.5 rounded-lg bg-purple-500/10 text-purple-500 w-fit">
            <Terminal className="w-5 h-5" />
          </div>
          <h4 className="font-bold text-sm">Declutr CLI Binary</h4>
          <p className="text-xs text-muted-foreground font-mono">declutr v1.0.0 (x86_64)</p>
          <button className="w-full py-1.5 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold flex items-center justify-center gap-1 transition-all">
            <Download className="w-3.5 h-3.5" /> Download CLI
          </button>
        </div>
      </div>
    </div>
  );
}
