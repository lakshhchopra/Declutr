"use client";

import React, { useEffect, useState } from "react";
import { Brain, FileText, Tag as TagIcon, CheckCircle2, AlertTriangle, MessageSquare, ListTodo, BarChart2 } from "lucide-react";
import { AnalysisService, AIAnalysis, Tag } from "../services/analysis-service";

export function AnalysisPanel({ assetId }: { assetId: string }) {
  const [analysis, setAnalysis] = useState<AIAnalysis | null>(null);

  useEffect(() => {
    AnalysisService.getAnalysis(assetId).then(setAnalysis);
  }, [assetId]);

  if (!analysis) {
    return (
      <div className="p-6 text-indigo-400 animate-pulse text-sm flex items-center gap-2">
        <Brain className="h-4 w-4 animate-pulse" /> Analyzing content...
      </div>
    );
  }

  const ConfidenceBadge = ({ score }: { score: number }) => {
    const color = score > 0.9 ? 'text-emerald-400' : score > 0.7 ? 'text-amber-400' : 'text-rose-400';
    return <span className={`text-[10px] font-mono ${color}`}>{Math.round(score * 100)}%</span>;
  };

  return (
    <div className="w-80 border-l border-slate-800 bg-slate-900 h-full overflow-y-auto hide-scrollbar">
      <div className="p-4 border-b border-slate-800 sticky top-0 bg-slate-900/95 backdrop-blur z-10 flex justify-between items-center">
        <h3 className="font-semibold text-white text-sm flex items-center gap-2">
          <Brain className="h-4 w-4 text-indigo-400" /> AI Understanding
        </h3>
        <ConfidenceBadge score={analysis.confidenceScore} />
      </div>

      <div className="p-4 space-y-6">
        {/* Summary Card */}
        <section>
          <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
            <FileText className="h-3 w-3" /> Summary
          </h4>
          <div className="bg-slate-800/50 border border-slate-700/50 rounded-lg p-3">
            <h5 className="text-sm font-semibold text-slate-200 mb-1">{analysis.title}</h5>
            <p className="text-xs text-slate-400 mb-3">{analysis.shortSummary}</p>
            <p className="text-xs text-slate-300 leading-relaxed pt-2 border-t border-slate-700/50">{analysis.detailedSummary}</p>
          </div>
        </section>

        {/* Classification Card */}
        <section>
          <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
            <ListTodo className="h-3 w-3" /> Classification
          </h4>
          <div className="space-y-2 text-sm bg-slate-800/30 p-3 rounded-lg border border-slate-800/80">
            <div className="flex justify-between items-center">
              <span className="text-slate-500">Category</span>
              <span className="text-slate-200">{analysis.classification.documentCategory}</span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-slate-500">Type</span>
              <span className="text-slate-200">{analysis.classification.documentType}</span>
            </div>
            <div className="flex justify-between items-center mt-2 pt-2 border-t border-slate-700/50">
              <span className="text-slate-500">Quality Score</span>
              <div className="flex items-center gap-2">
                <div className="w-16 h-1.5 bg-slate-700 rounded-full overflow-hidden">
                  <div className="h-full bg-emerald-500" style={{ width: `${analysis.classification.qualityScore * 100}%` }} />
                </div>
                <span className="text-slate-400 text-xs font-mono">{Math.round(analysis.classification.qualityScore * 100)}%</span>
              </div>
            </div>
            
            {/* Flags */}
            <div className="flex gap-2 mt-2">
              {analysis.classification.isScanned && <span className="px-2 py-0.5 bg-amber-500/10 text-amber-400 text-[10px] rounded border border-amber-500/20">Scanned</span>}
              {analysis.classification.isCorrupted && <span className="px-2 py-0.5 bg-rose-500/10 text-rose-400 text-[10px] rounded border border-rose-500/20">Corrupted</span>}
              {analysis.classification.isIncomplete && <span className="px-2 py-0.5 bg-amber-500/10 text-amber-400 text-[10px] rounded border border-amber-500/20">Incomplete</span>}
            </div>
          </div>
        </section>

        {/* Tags & Topics */}
        <section>
          <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
            <TagIcon className="h-3 w-3" /> Tags & Topics
          </h4>
          <div className="flex flex-wrap gap-2">
            {analysis.topics.map((t, i) => (
              <div key={`topic-${i}`} className="px-2 py-1 bg-indigo-500/10 border border-indigo-500/20 rounded flex items-center gap-1.5">
                <span className="text-indigo-300 text-xs">{t.name}</span>
                <div className="w-1.5 h-1.5 rounded-full bg-indigo-500" title={`Confidence: ${t.confidenceScore}`} />
              </div>
            ))}
            {analysis.tags.map((t, i) => (
              <div key={`tag-${i}`} className="px-2 py-1 bg-slate-800 border border-slate-700 rounded flex items-center gap-1.5">
                <span className="text-slate-300 text-xs">{t.name}</span>
                <div className="w-1.5 h-1.5 rounded-full bg-slate-500" title={`Confidence: ${t.confidenceScore}`} />
              </div>
            ))}
          </div>
        </section>

        {/* Metrics */}
        <section>
          <h4 className="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-3 flex items-center gap-2">
            <BarChart2 className="h-3 w-3" /> Analysis Metrics
          </h4>
          <div className="grid grid-cols-2 gap-2">
            <div className="bg-slate-800/30 p-2 rounded border border-slate-800/80">
              <span className="text-slate-500 text-[10px] uppercase block mb-1">Sentiment</span>
              <span className="text-slate-200 text-xs">{analysis.sentiment}</span>
            </div>
            <div className="bg-slate-800/30 p-2 rounded border border-slate-800/80">
              <span className="text-slate-500 text-[10px] uppercase block mb-1">Complexity</span>
              <span className="text-slate-200 text-xs">{analysis.complexity}</span>
            </div>
            <div className="bg-slate-800/30 p-2 rounded border border-slate-800/80">
              <span className="text-slate-500 text-[10px] uppercase block mb-1">Style</span>
              <span className="text-slate-200 text-xs truncate block">{analysis.writingStyle}</span>
            </div>
            <div className="bg-slate-800/30 p-2 rounded border border-slate-800/80">
              <span className="text-slate-500 text-[10px] uppercase block mb-1">Reading Level</span>
              <span className="text-slate-200 text-xs">{analysis.readingLevel}</span>
            </div>
          </div>
        </section>

      </div>
    </div>
  );
}
