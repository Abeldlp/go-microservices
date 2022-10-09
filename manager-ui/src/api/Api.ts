import axios, { AxiosResponse } from "axios";

class ApiBuilder {
  get(route: string): Promise<AxiosResponse> {
    return axios.get(route);
  }

  post<TData>(route: string, data: TData): Promise<AxiosResponse> {
    return axios.post(route, data);
  }

  put<TData>(route: string, data: TData): Promise<AxiosResponse> {
    return axios.put(route, data);
  }

  delete(route: string): Promise<AxiosResponse> {
    return axios.delete(route);
  }
}

const api = new ApiBuilder();

export { api, ApiBuilder };
