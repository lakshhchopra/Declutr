'use client';

import React, { useState, useEffect } from 'react';
import type { ShareComment } from '../types/collaboration';
import { CollaborationService } from '../services/collaboration-service';

export function CommentPanel() {
  const [comments, setComments] = useState<ShareComment[]>([]);
  const [input, setInput] = useState('');

  const loadComments = async () => {
    const list = await CollaborationService.getComments('share-japan-001');
    setComments(list);
  };

  useEffect(() => {
    loadComments();
  }, []);

  const handleAdd = async () => {
    if (!input.trim()) return;
    await CollaborationService.addComment('share-japan-001', input);
    setInput('');
    loadComments();
  };

  return (
    <div style={styles.container}>
      <h4 style={styles.title}>💬 Threaded Discussion Comments</h4>

      {/* Input */}
      <div style={styles.inputRow}>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Add a comment or mention team members..."
          style={styles.input}
          onKeyDown={(e) => e.key === 'Enter' && handleAdd()}
        />
        <button style={styles.btn} onClick={handleAdd}>
          Comment
        </button>
      </div>

      {/* Comments List */}
      <div style={styles.feed}>
        {comments.map((c) => (
          <div key={c.commentId} style={styles.card}>
            <div style={styles.header}>
              <span style={styles.author}>{c.userName}</span>
              <span style={styles.date}>{new Date(c.createdAt).toLocaleTimeString()}</span>
            </div>
            <p style={styles.body}>{c.content}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px', display: 'flex', flexDirection: 'column', gap: '14px' },
  title: { fontSize: '16px', fontWeight: 700, color: '#e2e8f0', margin: 0 },
  inputRow: { display: 'flex', gap: '10px' },
  input: { flex: 1, background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  btn: { background: 'linear-gradient(135deg, #6366f1, #818cf8)', color: '#fff', border: 'none', borderRadius: '10px', padding: '10px 18px', fontSize: '13px', fontWeight: 700, cursor: 'pointer' },
  feed: { display: 'flex', flexDirection: 'column', gap: '10px' },
  card: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '12px 14px', display: 'flex', flexDirection: 'column', gap: '4px' },
  header: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  author: { fontSize: '13px', fontWeight: 700, color: '#38bdf8' },
  date: { fontSize: '11px', color: '#64748b' },
  body: { fontSize: '13px', color: '#cbd5e1', margin: 0, lineHeight: 1.4 },
};
