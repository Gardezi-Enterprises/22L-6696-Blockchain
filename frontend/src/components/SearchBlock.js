import React, { useState } from 'react';

export default function SearchBlock({ api, onError }) {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState(null);
  const [loading, setLoading] = useState(false);
  const [searched, setSearched] = useState(false);

  const handleSearch = async () => {
    const q = query.trim();
    if (!q) {
      onError('Please enter a search query.');
      return;
    }
    setLoading(true);
    setResults(null);
    try {
      const res = await fetch(`${api}/search?q=${encodeURIComponent(q)}`);
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || 'Search failed');
      setResults(data.results || []);
      setSearched(true);
    } catch (err) {
      onError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const highlight = (text, query) => {
    if (!query) return text;
    const idx = text.toLowerCase().indexOf(query.toLowerCase());
    if (idx === -1) return text;
    return (
      <>
        {text.slice(0, idx)}
        <mark style={{ background: '#fef08a', color: '#1e293b', borderRadius: 2, padding: '0 2px' }}>
          {text.slice(idx, idx + query.length)}
        </mark>
        {text.slice(idx + query.length)}
      </>
    );
  };

  return (
    <div>
      <div className="card">
        <div className="card-title">🔍 Search Blockchain</div>
        <p style={{ color: 'var(--text-muted)', fontSize: '0.88rem', marginBottom: 20 }}>
          Search for any text across all transaction data in the blockchain. Matches are
          highlighted in the results.
        </p>

        <div className="search-bar">
          <input
            type="text"
            placeholder="Search transactions... e.g. Alice, Bob, BTC"
            value={query}
            onChange={e => setQuery(e.target.value)}
            onKeyDown={e => e.key === 'Enter' && handleSearch()}
          />
          <button
            className="btn btn-primary"
            onClick={handleSearch}
            disabled={loading || !query.trim()}
          >
            {loading ? '...' : '🔍 Search'}
          </button>
        </div>
      </div>

      {loading && (
        <div className="loading">
          <div className="loading-dots">
            <span /><span /><span />
          </div>
          <p style={{ marginTop: 12 }}>Searching blockchain...</p>
        </div>
      )}

      {!loading && searched && results !== null && (
        results.length === 0 ? (
          <div className="empty-state">
            <div className="icon">🔎</div>
            <p>No blocks found containing <strong>"{query}"</strong>.</p>
          </div>
        ) : (
          <div>
            <p className="result-count">
              Found <strong style={{ color: 'var(--accent-hover)' }}>{results.length}</strong> block(s) matching "{query}"
            </p>
            {results.map(block => (
              <div key={block.index} className="block-card" style={{ marginBottom: 16 }}>
                <div className="block-header" style={{ cursor: 'default' }}>
                  <div className="block-header-left">
                    <span className="block-index">#{block.index}</span>
                    <div>
                      <div className="block-title">
                        {block.index === 0 ? '🌱 Genesis Block' : `Block #${block.index}`}
                      </div>
                      <div className="block-subtitle">
                        {block.timestamp} · Nonce: {block.nonce}
                      </div>
                    </div>
                  </div>
                  <span style={{ background: 'var(--success)', color: '#fff', padding: '3px 10px', borderRadius: 6, fontSize: '0.75rem', fontWeight: 700 }}>
                    MATCH
                  </span>
                </div>
                <div className="block-body">
                  <div className="block-field">
                    <span className="field-key">Hash</span>
                    <span className="field-value hash">{block.hash}</span>
                  </div>
                  <div className="transactions-list">
                    <div className="transactions-list-title">Transactions</div>
                    {(block.data || []).map((tx, i) => {
                      const isMatch = tx.toLowerCase().includes(query.toLowerCase());
                      return (
                        <div
                          key={i}
                          className="tx-item"
                          style={isMatch ? { borderColor: 'var(--warning)', background: '#1c1a0e' } : {}}
                        >
                          {isMatch && <span style={{ color: 'var(--warning)', marginRight: 8 }}>★</span>}
                          {highlight(tx, query)}
                        </div>
                      );
                    })}
                  </div>
                </div>
              </div>
            ))}
          </div>
        )
      )}
    </div>
  );
}
