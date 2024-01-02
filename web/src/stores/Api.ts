/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ResponseError {
  message?: string;
}

export interface ResponseOk {
  message?: string;
}

export interface ModelsDatarefValue {
  name?: string;
  value?: any;
}

export interface ModelsFlightInfo {
  airportId?: string;
  airportName?: string;
  fuelWeight?: number;
  time?: number;
  totalWeight?: number;
}

export enum ModelsFlightState {
  FlightStateParked = "parked",
  FlightStateTaxiOut = "taxi_out",
  FlightStateTakeoff = "takeoff",
  FlightStateClimb = "climb",
  FlightStateCruise = "cruise",
  FlightStateDescend = "descend",
  FlightStateLanding = "landing",
  FlightStateTaxiIn = "taxi_in",
}

export interface ModelsFlightStatus {
  aircraftDisplayName?: string;
  aircraftICAO?: string;
  arrivalFlightInfo?: ModelsFlightInfo;
  createdAt?: string;
  departureFlightInfo?: ModelsFlightInfo;
  events?: ModelsFlightStatusEvent[];
  id?: number;
  locations?: ModelsFlightStatusLocation[];
  updatedAt?: string;
  username?: string;
}

export interface ModelsFlightStatusEvent {
  createdAt?: string;
  description?: string;
  details?: string;
  eventType?: ModelsFlightStatusEventType;
  flightId?: number;
  id?: number;
  timestamp?: number;
  updatedAt?: string;
}

export enum ModelsFlightStatusEventType {
  StateEvent = "event:state",
  LocationEvent = "event:location",
  ViolationEvent = "event:violation",
}

export interface ModelsFlightStatusLocation {
  agl?: number;
  altitude?: number;
  createdAt?: string;
  flapRatio?: number;
  flightId?: number;
  fuel?: number;
  gearForce?: number;
  gforce?: number;
  gs?: number;
  heading?: number;
  ias?: number;
  id?: number;
  lat?: number;
  lng?: number;
  pitch?: number;
  state?: ModelsFlightState;
  timestamp?: number;
  updatedAt?: string;
  vs?: number;
}

export interface ModelsSendCommandReq {
  command?: string;
}

export interface ModelsSetDatarefValue {
  dataref?: string;
  value?: any;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType | null) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
  Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "/apis";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) => fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(typeof value === "number" ? value : `${value}`)}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) => (Array.isArray(query[key]) ? this.addArrayQueryParam(query, key) : this.addQueryParam(query, key)))
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.Text]: (input: any) => (input !== null && typeof input !== "string" ? JSON.stringify(input) : input),
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === "object" && property !== null
            ? JSON.stringify(property)
            : `${property}`,
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(requestParams.headers || {}),
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
      },
      signal: (cancelToken ? this.createAbortSignal(cancelToken) : requestParams.signal) || null,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title No title
 * @baseUrl /apis
 * @contact
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  flightLogs = {
    /**
     * No description
     *
     * @tags Flight_Logs
     * @name FlightLogsList
     * @summary Get a list of FlightLogs
     * @request GET:/flight-logs
     */
    flightLogsList: (
      query?: {
        /** specify if it's overview */
        isOverview?: string;
        /** departure airport */
        departureAirportId?: string;
        /** arrival airport */
        arrivalAirportId?: string;
        /** aircraft ICAO */
        aircraftICAO?: string;
        /** xplane or xws */
        source?: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<ModelsFlightStatus[], ResponseError>({
        path: `/flight-logs`,
        method: "GET",
        query: query,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Flight_Logs
     * @name FlightLogsDetail
     * @summary Get one FlightLog
     * @request GET:/flight-logs/{id}
     */
    flightLogsDetail: (id: string, params: RequestParams = {}) =>
      this.request<ModelsFlightStatus, void>({
        path: `/flight-logs/${id}`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  liveness = {
    /**
     * No description
     *
     * @tags Misc
     * @name LivenessList
     * @summary K8s Liveness endpoint
     * @request GET:/liveness
     */
    livenessList: (params: RequestParams = {}) =>
      this.request<void, any>({
        path: `/liveness`,
        method: "GET",
        type: ContentType.Json,
        ...params,
      }),
  };
  readiness = {
    /**
     * No description
     *
     * @tags Misc
     * @name ReadinessList
     * @summary K8s Readiness endpoint
     * @request GET:/readiness
     */
    readinessList: (params: RequestParams = {}) =>
      this.request<void, ResponseError>({
        path: `/readiness`,
        method: "GET",
        type: ContentType.Json,
        ...params,
      }),
  };
  version = {
    /**
     * No description
     *
     * @tags Misc
     * @name VersionList
     * @summary Get version of GPT X-Plane plugin
     * @request GET:/version
     */
    versionList: (params: RequestParams = {}) =>
      this.request<ResponseOk, ResponseError>({
        path: `/version`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  xplm = {
    /**
     * No description
     *
     * @tags Dataref
     * @name CommandUpdate
     * @summary Send command to X Plane
     * @request PUT:/xplm/command
     * @secure
     */
    commandUpdate: (request: ModelsSendCommandReq, params: RequestParams = {}) =>
      this.request<any, void>({
        path: `/xplm/command`,
        method: "PUT",
        body: request,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * No description
     *
     * @tags Dataref
     * @name DatarefList
     * @summary Get Dataref
     * @request GET:/xplm/dataref
     * @secure
     */
    datarefList: (
      query: {
        /** xplane dataref string */
        dataref_str: string;
        /** alias name, if not set, dataref_str will be used */
        alias?: string;
        /** -1: raw, 2: round up to two digits */
        precision: number;
        /** transform xplane byte array to string */
        is_byte_array?: boolean;
      },
      params: RequestParams = {},
    ) =>
      this.request<ModelsDatarefValue, ResponseError>({
        path: `/xplm/dataref`,
        method: "GET",
        query: query,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Dataref
     * @name DatarefUpdate
     * @summary Set Dataref
     * @request PUT:/xplm/dataref
     * @secure
     */
    datarefUpdate: (request: ModelsSetDatarefValue, params: RequestParams = {}) =>
      this.request<any, void>({
        path: `/xplm/dataref`,
        method: "PUT",
        body: request,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * No description
     *
     * @tags Dataref
     * @name DatarefsUpdate
     * @summary Set a list of Dataref
     * @request PUT:/xplm/datarefs
     * @secure
     */
    datarefsUpdate: (params: RequestParams = {}) =>
      this.request<any, void>({
        path: `/xplm/datarefs`,
        method: "PUT",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * No description
     *
     * @tags Dataref
     * @name DatarefsCreate
     * @summary Get a list of Dataref
     * @request POST:/xplm/datarefs
     * @secure
     */
    datarefsCreate: (params: RequestParams = {}) =>
      this.request<ModelsDatarefValue[], void>({
        path: `/xplm/datarefs`,
        method: "POST",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
}
