export const DEFAULT_OG_IMAGE = "/og-image.png";

export const getSiteUrl = (): string => {
  return import.meta.env.VITE_SITE_URL ?? "http://localhost:3000";
};

export const getOGImage = () => {
  return `${getSiteUrl()}${DEFAULT_OG_IMAGE}`;
};
