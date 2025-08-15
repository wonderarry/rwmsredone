import { format, formatDistanceToNowStrict, parseISO } from 'date-fns';

export function formatDate(input: Date | string, pattern = 'yyyy-MM-dd HH:mm') {
  const d = typeof input === 'string' ? parseISO(input) : input;
  return format(d, pattern);
}

export function timeAgo(input: Date | string) {
  const d = typeof input === 'string' ? parseISO(input) : input;
  return formatDistanceToNowStrict(d, { addSuffix: true });
}

export function shortId(id: string, head = 4, tail = 4) {
  if (!id) return '';
  return id.length <= head + tail ? id : `${id.slice(0, head)}…${id.slice(-tail)}`;
}

export function truncate(text: string, max = 120) {
  if (text.length <= max) return text;
  return text.slice(0, Math.max(0, max - 1)).trimEnd() + '…';
}