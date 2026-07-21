"use client";

import React, { useEffect, useState } from "react";
import { Info, Camera, FileText, Database, Activity, MapPin } from "lucide-react";
import { Card, CardContent } from "../../../shared/components/ui/card";
import { MetadataService, CompleteMetadata } from "../services/metadata-service";
import { formatBytes } from "../../../shared/utils/format";

export function MetadataPanel({ assetId }: { assetId: string }) {
  const [metadata, setMetadata] = useState<CompleteMetadata | null>(null);

  useEffect(() => {
    MetadataService.getMetadata(assetId).then(setMetadata);
  }, [assetId]);

  if (!metadata) {
    return <div className="p-6 text-slate-400 animate-pulse text-sm">Loading metadata...</div>;
  }

  const { general, properties, exif } = metadata;

  return (
    <div className="w-80 border-l border-slate-800 bg-slate-900 h-full overflow-y-auto hide-scrollbar">
      <div className="p-4 border-b border-slate-800 sticky top-0 bg-slate-900/95 backdrop-blur z-10 flex justify-between items-center">
        <h3 className="font-semibold text-white text-sm flex items-center gap-2">
          <Info className="h-4 w-4 text-emerald-400" /> Information
        </h3>
      </div>

      <div className="p-4 space-y-6">
        {/* General */}
        <section>
          <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
            <FileText className="h-3 w-3" /> General
          </h4>
          <div className="space-y-2 text-sm">
            <div className="flex justify-between">
              <span className="text-slate-500">Name</span>
              <span className="text-slate-200 truncate max-w-[150px]" title={general.filename}>{general.filename}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-slate-500">Type</span>
              <span className="text-slate-200">{general.mimeType}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-slate-500">Size</span>
              <span className="text-slate-200">{formatBytes(general.fileSize)}</span>
            </div>
          </div>
        </section>

        {/* Properties */}
        {properties && Object.keys(properties.properties).length > 0 && (
          <section>
            <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
              <Database className="h-3 w-3" /> Technical
            </h4>
            <div className="space-y-2 text-sm">
              {Object.entries(properties.properties).map(([key, value]) => (
                <div key={key} className="flex justify-between">
                  <span className="text-slate-500 capitalize">{key.replace(/([A-Z])/g, ' $1').trim()}</span>
                  <span className="text-slate-200">{String(value)}</span>
                </div>
              ))}
            </div>
          </section>
        )}

        {/* EXIF */}
        {exif && exif.cameraMake && (
          <section>
            <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
              <Camera className="h-3 w-3" /> EXIF Data
            </h4>
            <div className="space-y-2 text-sm bg-slate-800/50 p-3 rounded-lg border border-slate-800">
              {exif.cameraMake && (
                <div className="flex justify-between">
                  <span className="text-slate-500">Camera</span>
                  <span className="text-slate-200 text-right">{exif.cameraMake} {exif.cameraModel}</span>
                </div>
              )}
              {exif.lens && (
                <div className="flex justify-between mt-1">
                  <span className="text-slate-500">Lens</span>
                  <span className="text-slate-200 text-right truncate max-w-[120px]" title={exif.lens}>{exif.lens}</span>
                </div>
              )}
              <div className="flex justify-between mt-2 pt-2 border-t border-slate-700/50">
                <span className="text-slate-500">Settings</span>
                <span className="text-slate-200 font-mono text-xs">
                  ƒ/{exif.fStop}  {exif.exposure}s  ISO{exif.iso}
                </span>
              </div>
              {exif.gpsLat && exif.gpsLong && (
                <div className="flex justify-between mt-2 pt-2 border-t border-slate-700/50">
                  <span className="text-slate-500 flex items-center gap-1"><MapPin className="h-3 w-3" /> Location</span>
                  <span className="text-blue-400 cursor-pointer text-xs">
                    {exif.gpsLat.toFixed(4)}, {exif.gpsLong.toFixed(4)}
                  </span>
                </div>
              )}
            </div>
          </section>
        )}
        
        {/* Tracking */}
        <section>
          <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
            <Activity className="h-3 w-3" /> System
          </h4>
          <div className="space-y-2 text-xs">
            <div className="flex justify-between">
              <span className="text-slate-500">SHA-256</span>
              <span className="text-slate-400 font-mono truncate max-w-[140px]" title={general.hash}>{general.hash}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-slate-500">Extracted</span>
              <span className="text-slate-400">{new Date(general.lastExtractedAt).toLocaleString()}</span>
            </div>
          </div>
        </section>

      </div>
    </div>
  );
}
