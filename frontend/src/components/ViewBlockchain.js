import React, { useState } from 'react';

export default function ViewBlockchain({ chain, loading }) {
  const [expandedIdx, setExpandedIdx] = useState(null);

  if (loading) {
    return (
      <div className="loading">
        <div className="loading-dots">
          <span /><span /><span />
        </div>
        <p style={{ marginTop: 12 }}>Loading blockchain...</p>
      </div>
    );
  }

  if (!chain || chain.length === 0) {
    return (
      <div className="empty-state">
        <div className="icon">⛓</div>
        <p>No blocks found. Make sure the Go server is running.</p>
      </div>
    );
  }

  const toggleBlock = (idx) => {
    setExpandedIdx(expandedIdx === idx ? null : idx);
  };

  const truncateHash = (hash) => {
    if (!hash) return '—';
    return hash.length > 20 ? `${hash.slice(0, 10)}...${hash.slice(-10)}` : hash;
  };

  return (
    <div>
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 20 }}>
        <h2 style={{ fontSize: '1.1rem', fontWeight: 600 }}>
          🔗 Blockchain — <span style={{ color: 'var(--accent-hover)' }}>{chain.length} block(s)</span>
        </h2>
        <span className="merkle-badge">🌿 Merkle Tree Enabled</span>
      </div>

      <div className="chain-container">
        {chain.map((block, i) => {
          const isOpen = expandedIdx === block.index;
          const isGenesis = block.index === 0;
          return (
            <React.Fragment key={block.index}>
              <div className={`block-card ${isGenesis ? 'genesis' : ''}`}>
                <div className="block-header" onClick={() => toggleBlock(block.index)}>
                  <div className="block-header-left">
                    <span className="block-index">#{block.index}</span>
                    <div>
                      <div className="block-title">
                        {isGenesis ? '🌱 Genesis Block' : `Block #${block.index}`}
                      </div>
                      <div className="block-subtitle">
                        {block.timestamp} · {block.data?.length || 0} transaction(s) · Nonce: {block.nonce}
                      </div>
                    </div>
                  </div>
                  <span className={`block-toggle ${isOpen ? 'open' : ''}`}>▼</span>
                </div>

                {isOpen && (
                  <div className="block-body">
                    <div className="block-field">
                      <span className="field-key">Index</span>
                      <span className="field-value">{block.index}</span>
                    </div>
                    <div className="block-field">
                      <span className="field-key">Timestamp</span>
                      <span className="field-value">{block.timestamp}</span>
                    </div>
                    <div className="block-field">
                      <span className="field-key">Difficulty</span>
                      <span className="field-value">{block.difficulty}</span>
                    </div>
                    <div className="block-field">
                      <span className="field-key">Nonce</span>
                      <span className="field-value">{block.nonce}</span>
                    </div>
                    <div className="block-field">
                      <span className="field-key">Merkle Root</span>
                      <span className="field-value hash" title={block.merkleRoot}>
                        {truncateHash(block.merkleRoot)}
                      </span>
                    </div>
                    <div className="block-field">
                      <span className="field-key">Prev Hash</span>
                      <span className="field-value hash" title={block.prevHash}>
                        {truncateHash(block.prevHash)}
                      </span>
                    </div>
                    <div className="block-field">
                      <span className="field-key">Hash</span>
                      <span className="field-value hash good" title={block.hash}>
                        {block.hash}
                      </span>
                    </div>

                    <div className="transactions-list">
                      <div className="transactions-list-title">
                        Transactions ({block.data?.length || 0})
                      </div>
                      {(block.data || []).map((tx, ti) => (
                        <div key={ti} className="tx-item">
                          <span style={{ color: 'var(--text-muted)', marginRight: 8, fontFamily: 'var(--mono)', fontSize: '0.78rem' }}>
                            #{ti + 1}
                          </span>
                          {tx}
                        </div>
                      ))}
                    </div>
                  </div>
                )}
              </div>

              {i < chain.length - 1 && (
                <div className="chain-connector">↕</div>
              )}
            </React.Fragment>
          );
        })}
      </div>
    </div>
  );
}
