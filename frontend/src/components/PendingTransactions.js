import React from 'react';

export default function PendingTransactions({ pending, onRefresh }) {
  return (
    <div>
      <div className="card">
        <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 18 }}>
          <div className="card-title" style={{ margin: 0 }}>
            ⏳ Pending Transactions
            {pending.length > 0 && (
              <span style={{
                background: 'var(--warning)',
                color: '#000',
                borderRadius: 12,
                padding: '2px 8px',
                fontSize: '0.75rem',
                fontWeight: 700,
                marginLeft: 8,
              }}>
                {pending.length}
              </span>
            )}
          </div>
          <button
            className="refresh-btn"
            onClick={onRefresh}
            title="Refresh pending transactions"
          >
            ↻ Refresh
          </button>
        </div>

        <p style={{ color: 'var(--text-muted)', fontSize: '0.88rem', marginBottom: 20 }}>
          These transactions are waiting to be included in the next mined block.
          Click <strong>⛏ Mine Block</strong> to mine them into the chain.
        </p>

        {pending.length === 0 ? (
          <div className="empty-state">
            <div className="icon">📭</div>
            <p>No pending transactions.</p>
            <p style={{ marginTop: 8, fontSize: '0.82rem' }}>
              Go to <strong>➕ Add Transaction</strong> to add some.
            </p>
          </div>
        ) : (
          <div className="pending-list">
            {pending.map((tx, i) => (
              <div key={i} className="pending-item">
                <span className="pending-num">#{i + 1}</span>
                <span>{tx}</span>
              </div>
            ))}
          </div>
        )}
      </div>

      {pending.length > 0 && (
        <div className="card" style={{ borderColor: '#f59e0b40' }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
            <span style={{ fontSize: '1.4rem' }}>ℹ️</span>
            <div>
              <p style={{ fontWeight: 600, fontSize: '0.9rem' }}>Ready to mine</p>
              <p style={{ color: 'var(--text-muted)', fontSize: '0.85rem', marginTop: 4 }}>
                You have {pending.length} transaction(s) waiting. Switch to the{' '}
                <strong style={{ color: 'var(--warning)' }}>⛏ Mine Block</strong> tab to bundle
                them into a block using Proof of Work.
              </p>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
