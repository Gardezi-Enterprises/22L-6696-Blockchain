import React, { useState } from 'react';

export default function MineBlock({ api, pendingCount, onSuccess, onError }) {
  const [mining, setMining] = useState(false);
  const [lastMined, setLastMined] = useState(null);
  const [elapsed, setElapsed] = useState(null);

  const handleMine = async () => {
    if (pendingCount === 0) {
      onError('No pending transactions. Add at least one transaction first.');
      return;
    }
    setMining(true);
    setLastMined(null);
    const start = Date.now();

    try {
      const res = await fetch(`${api}/mine`, { method: 'POST' });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || 'Mining failed');
      const took = ((Date.now() - start) / 1000).toFixed(2);
      setLastMined(data.block);
      setElapsed(took);
      onSuccess(`Block #${data.block.index} mined in ${took}s!`);
    } catch (err) {
      onError(err.message);
    } finally {
      setMining(false);
    }
  };

  return (
    <div>
      <div className="card">
        <div className="card-title">⛏ Mine a New Block</div>
        <p style={{ color: 'var(--text-muted)', fontSize: '0.88rem', marginBottom: 20 }}>
          Mining runs Proof of Work — the server iterates nonces until it finds a hash with
          the required number of leading zeros (difficulty). All pending transactions are
          included in the new block and a Merkle Root is computed.
        </p>

        <div className="mine-info">
          <div className="mine-info-box">
            <div className="mine-info-value">{pendingCount}</div>
            <div className="mine-info-label">Pending Transactions</div>
          </div>
          <div className="mine-info-box">
            <div className="mine-info-value" style={{ color: 'var(--accent-hover)' }}>3</div>
            <div className="mine-info-label">Difficulty (leading zeros)</div>
          </div>
          <div className="mine-info-box">
            <div className="mine-info-value" style={{ color: 'var(--success)' }}>SHA-256</div>
            <div className="mine-info-label">Hash Algorithm</div>
          </div>
        </div>

        <div className="mine-panel">
          <button
            className="btn btn-mine"
            onClick={handleMine}
            disabled={mining}
          >
            {mining ? (
              <>
                <span className="mining-spinner" />
                Mining in Progress...
              </>
            ) : '⛏ Mine Block'}
          </button>

          {mining && (
            <p style={{ marginTop: 16, color: 'var(--warning)', fontSize: '0.88rem', animation: 'pulse 1s infinite' }}>
              Searching for a valid nonce... this may take a moment.
            </p>
          )}

          {pendingCount === 0 && !mining && (
            <p style={{ marginTop: 14, color: 'var(--text-muted)', fontSize: '0.85rem' }}>
              ⚠️ Add transactions first before mining.
            </p>
          )}
        </div>
      </div>

      {lastMined && (
        <div className="card">
          <div className="card-title" style={{ color: 'var(--success)' }}>
            ✅ Block #{lastMined.index} Mined in {elapsed}s
          </div>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 10 }}>
            {[
              { label: 'Block Index',   value: lastMined.index },
              { label: 'Timestamp',     value: lastMined.timestamp },
              { label: 'Nonce',         value: lastMined.nonce },
              { label: 'Difficulty',    value: lastMined.difficulty },
              { label: 'Transactions',  value: lastMined.data?.length },
              { label: 'Merkle Root',   value: lastMined.merkleRoot },
              { label: 'Hash',          value: lastMined.hash },
            ].map(f => (
              <div key={f.label} className="block-field">
                <span className="field-key">{f.label}</span>
                <span className="field-value hash">{f.value}</span>
              </div>
            ))}
          </div>

          <div className="transactions-list">
            <div className="transactions-list-title">Included Transactions</div>
            {(lastMined.data || []).map((tx, i) => (
              <div key={i} className="tx-item">
                <span style={{ color: 'var(--text-muted)', marginRight: 8, fontFamily: 'var(--mono)', fontSize: '0.78rem' }}>
                  #{i + 1}
                </span>
                {tx}
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
