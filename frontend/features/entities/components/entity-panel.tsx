"use client";

import React, { useEffect, useState } from "react";
import { Link2, Building2, MapPin, Calendar, DollarSign, Package, Fingerprint, User, Hash } from "lucide-react";
import { EntityService, AssetEntityResponse, EntityType } from "../services/entity-service";

export function EntityPanel({ assetId }: { assetId: string }) {
  const [entities, setEntities] = useState<AssetEntityResponse[] | null>(null);

  useEffect(() => {
    EntityService.getEntitiesForAsset(assetId).then(setEntities);
  }, [assetId]);

  if (!entities) {
    return (
      <div className="p-6 text-indigo-400 animate-pulse text-sm flex items-center gap-2">
        <Link2 className="h-4 w-4 animate-pulse" /> Extracting Entities...
      </div>
    );
  }

  const getIconForType = (type: EntityType) => {
    switch (type) {
      case 'Organization': return <Building2 className="h-3 w-3" />;
      case 'Location': return <MapPin className="h-3 w-3" />;
      case 'Date': return <Calendar className="h-3 w-3" />;
      case 'Amount': return <DollarSign className="h-3 w-3" />;
      case 'Product': return <Package className="h-3 w-3" />;
      case 'Person': return <User className="h-3 w-3" />;
      case 'Identifier': return <Fingerprint className="h-3 w-3" />;
      default: return <Hash className="h-3 w-3" />;
    }
  };

  const getColorForType = (type: EntityType) => {
    switch (type) {
      case 'Organization': return 'bg-blue-500/10 text-blue-400 border-blue-500/20';
      case 'Location': return 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20';
      case 'Date': return 'bg-purple-500/10 text-purple-400 border-purple-500/20';
      case 'Amount': return 'bg-amber-500/10 text-amber-400 border-amber-500/20';
      case 'Product': return 'bg-rose-500/10 text-rose-400 border-rose-500/20';
      case 'Person': return 'bg-cyan-500/10 text-cyan-400 border-cyan-500/20';
      default: return 'bg-slate-500/10 text-slate-400 border-slate-500/20';
    }
  };

  // Group entities by type
  const grouped = entities.reduce((acc, curr) => {
    if (!acc[curr.entity.type]) acc[curr.entity.type] = [];
    acc[curr.entity.type].push(curr);
    return acc;
  }, {} as Record<string, AssetEntityResponse[]>);

  return (
    <div className="w-80 border-l border-slate-800 bg-slate-900 h-full overflow-y-auto hide-scrollbar flex flex-col">
      <div className="p-4 border-b border-slate-800 sticky top-0 bg-slate-900/95 backdrop-blur z-10">
        <h3 className="font-semibold text-white text-sm flex items-center gap-2">
          <Link2 className="h-4 w-4 text-indigo-400" /> Extracted Entities
        </h3>
        <p className="text-xs text-slate-500 mt-1">Canonical knowledge graph mapping.</p>
      </div>

      <div className="p-4 space-y-6 flex-1">
        {Object.entries(grouped).map(([type, items]) => (
          <section key={type}>
            <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
              {getIconForType(type as EntityType)} {type}
            </h4>
            <div className="flex flex-col gap-2">
              {items.map((item, idx) => (
                <div key={`${type}-${idx}`} className="group relative bg-slate-800/40 hover:bg-slate-800/80 transition-colors border border-slate-800/80 rounded-lg p-3 cursor-pointer">
                  
                  {/* Entity Chip */}
                  <div className="flex items-start justify-between">
                    <div>
                      <div className="flex items-center gap-2 mb-1">
                        <span className={`px-1.5 py-0.5 rounded text-[10px] uppercase font-bold border ${getColorForType(item.entity.type)}`}>
                          {item.entity.type}
                        </span>
                        <span className="text-slate-200 font-medium text-sm">{item.entity.canonicalName}</span>
                      </div>
                      <div className="text-xs text-slate-500 flex items-center gap-1">
                        Found as: <span className="font-mono text-slate-400">"{item.occurrence.originalValue}"</span>
                      </div>
                    </div>
                    <div className="w-1.5 h-1.5 rounded-full bg-emerald-500 mt-1" title={`Confidence: ${item.occurrence.confidenceScore}`} />
                  </div>

                  {/* Inspector Hover State */}
                  <div className="hidden group-hover:block mt-3 pt-3 border-t border-slate-700/50">
                    <p className="text-xs text-slate-400 mb-1">{item.entity.description || "No description available."}</p>
                    {item.entity.aliases.length > 0 && (
                      <div className="text-[10px] text-slate-500 flex gap-1 flex-wrap mt-2">
                        Aliases: 
                        {item.entity.aliases.map(a => (
                          <span key={a} className="bg-slate-900 px-1 rounded">{a}</span>
                        ))}
                      </div>
                    )}
                  </div>

                </div>
              ))}
            </div>
          </section>
        ))}
        
        {entities.length === 0 && (
          <div className="text-center py-10 text-slate-500 text-sm">
            No entities detected in this asset.
          </div>
        )}
      </div>
    </div>
  );
}
