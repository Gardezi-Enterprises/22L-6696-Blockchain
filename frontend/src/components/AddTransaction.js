import React, { useState } from 'react';

export default function AddTransaction({ api, onSuccess, onError }) {
  const [transaction, setTransaction] = useState('');
  const [loading, setLoading] = useState(false);
  const [history, setHistory] = useState([]);

  const handleAdd = async () => {
    const tx = transaction.trim();
    if (!tx) {
      onError('Transaction cannot be empty.');
      return;
    }
    setLoading(true);
    try {
      const res = await fetch(`${api}/transaction`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ transaction: tx }),
      });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || 'Failed to add transaction');
      setHistory(prev => [tx, ...prev]);
      setTransaction('');
      onSuccess(`Transaction added: "${tx}"`);
    } catch (err) {
      onError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const examples = [
    'Alice → Bob: 50 BTC',
    'Bob → Charlie: 25 ETH',
    'Charlie → Dave: 100 USDT',
    'Dave → Alice: 10 BNB',
  ];

  return (
    <div>
      <div className="card">
        <div className="card-title">➕ Add Transaction</div>
        <p style={{ color: 'var(--text-muted)', fontSize: '0.88rem', marginBottom: 20 }}>
          Transactions are stored in the pending pool and included in the next mined block.
          Each block computes a Merkle Root over its transactions.
        </p>

        <div className="form-group">
          <label>Transaction Data</label>
          <textarea
            placeholder="e.g. Alice → Bob: 50 BTC"
            value={transaction}
            onChange={e => setTransaction(e.target.value)}
            onKeyDown={e => e.ctrlKey && e.key === 'Enter' && handleAdd()}
          />
          <span style={{ fontSize: '0.75rem', color: 'var(--text-muted)' }}>Ctrl+Enter to submit</span>
        </div>

        <div style={{ marginBottom: 20 }}>
          <p style={{ fontSize: '0.78rem', color: 'var(--text-muted)', marginBottom: 8, textTransform: 'uppercase', letterSpacing: '0.05em', fontWeight: 600 }}>
            Quick Examples
          </p>
          <div style={{ display: 'flex', flexWrap: 'wrap', gap: 8 }}>
            {examples.map(ex => (
              <button
                key={ex}
                onClick={() => setTransaction(ex)}
                style={{
                  background: 'var(--surface2)',
                  border: '1px solid var(--border)',
                  color: 'var(--text-muted)',
                  padding: '5px 12px',
                  borderRadius: 6,
                  fontSize: '0.8rem',
                  cursor: 'pointer',
                  transition: 'all 0.2s',
                }}
                onMouseEnter={e => { e.target.style.color = 'var(--text)'; e.target.style.borderColor = 'var(--accent)'; }}
                onMouseLeave={e => { e.target.style.color = 'var(--text-muted)'; e.target.style.borderColor = 'var(--border)'; }}
              >
                {ex}
              </button>
            ))}
          </div>
        </div>

        <button className="btn btn-primary" onClick={handleAdd} disabled={loading || !transaction.trim()}>
          {loading ? (
            <>
              <span className="mining-spinner" style={{ width: 16, height: 16, borderWidth: 2 }} />
              Adding...
            </>
          ) : '➕ Add to Pending Pool'}
        </button>
      </div>

      {history.length > 0 && (
        <div className="card">
          <div className="card-title">📋 Added This Session ({history.length})</div>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
            {history.map((tx, i) => (
              <div key={i} className="tx-item">
                <span style={{ color: 'var(--success)', marginRight: 8 }}>✓</span>
                {tx}
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
