import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export const RWMS_LOGIN_TOKEN_KEY = 'rwms_token'

export default async function Home() {
    const c  = await cookies()
    const token = c.get(RWMS_LOGIN_TOKEN_KEY)?.value;
  redirect(token ? '/dashboard' : '/auth/login');
}