// src/shared/api/socket.ts
import { useUserStore } from '@/store';
import { WS_BASE_URL } from '../config';
import { refreshAccessToken } from './tokenService';
import { handleAuthError } from './authError';

type MessageHandler = (data: any) => void;

class SocketClient {
    private socket: WebSocket | null = null;
    private activePath: string | null = null;
    private listeners = new Map<string, Set<MessageHandler>>();
    private retryCount = 0;
    private maxRetries = 10;
    private retryDelay = 2000;
    private retryDelayMax = 10000;

    private manuallyClosed = false;
    /**
     * path: must begin with "/", e.g. "/ws/lounges/…"
     */
    async connect(path: string) {
        this.manuallyClosed = false;
        this.activePath = path;

        // close existing
        this.disconnect();

        // ensure token is fresh
        const userStore = useUserStore();
        let token = userStore.getToken;
        if (!token) {
            // try refresh once
            try {
                token = await refreshAccessToken();
            } catch {
                handleAuthError({ code: 'token_missing' });
                return;
            }
        }

        const url = `${WS_BASE_URL}${path}?token=${token}`;
        this.socket = new WebSocket(url);

        this.socket.onopen = () => {
            console.debug('[Socket] connected to', url);
        };

        this.socket.onmessage = (ev) => {
            let msg: any;
            try {
                msg = JSON.parse(ev.data);
                console.log('chat msg', msg);
            } catch {
                return;
            }
            const handlers = this.listeners.get(msg.type) ?? new Set();
            handlers.forEach((fn) => fn(msg));
        };

        this.socket.onclose = (ev) => {
            if (this.manuallyClosed || this.activePath !== path) {
                return;
            }
            console.warn('[Socket] closed, will retry in 2s', ev);
            this.retryCount++;
            if (this.retryCount < this.maxRetries) {
                const delay = Math.min(
                    this.retryCount * this.retryDelay,
                    this.retryDelayMax,
                );
                setTimeout(() => this.connect(path), delay);
            }
        };

        this.socket.onerror = (err) => {
            console.error('[Socket] error', err);
            this.socket?.close();
        };
    }

    disconnect() {
        if (this.socket) {
            this.manuallyClosed = true;
            this.socket.onclose = null;
            this.socket.close();
            this.socket = null;
        }
    }

    send(type: string, payload: any) {
        if (this.socket?.readyState === WebSocket.OPEN) {
            this.socket.send(JSON.stringify({ type, ...payload }));
        }
    }

    on(type: string, handler: MessageHandler) {
        const set = this.listeners.get(type) ?? new Set();
        set.add(handler);
        this.listeners.set(type, set);
    }

    off(type: string, handler: MessageHandler) {
        const set = this.listeners.get(type);
        if (set) {
            set.delete(handler);
            if (set.size === 0) this.listeners.delete(type);
        }
    }
}

export const socketClient = new SocketClient();
