package event

// NotificationEventData
type NotificationEventData struct {
	// Object containing the API resource relevant to the event. For example, an `invoice.created` event will have a full [invoice object](https://stripe.com/docs/api#invoice_object) as the value of the object key.
	Object map[string]interface{} `json:"object"`
	// Object containing the names of the attributes that have changed, and their previous values (sent along only with *.updated events).
	PreviousAttributes map[string]interface{} `json:"previous_attributes,omitempty"`
}

// NotificationEventRequest
type NotificationEventRequest struct {
	// ID of the API request that caused the event. If null, the event was automatic (e.g., Stripe's automatic subscription handling). Request logs are available in the [dashboard](https://dashboard.stripe.com/logs), but currently not in the API.
	Id string `json:"id,omitempty"`
	// The idempotency key transmitted during the request, if any. *Note: This property is populated only for events on or after May 23, 2017*.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// Event Events are our way of letting you know when something interesting happens in your account. When an interesting event occurs, we create a new `Event` object. For example, when a charge succeeds, we create a `charge.succeeded` event; and when an invoice payment attempt fails, we create an `invoice.payment_failed` event. Note that many API requests may cause multiple events to be created. For example, if you create a new subscription for a customer, you will receive both a `customer.subscription.created` event and a `charge.succeeded` event.  Events occur when the state of another API resource changes. The state of that resource at the time of the change is embedded in the event's data field. For example, a `charge.succeeded` event will contain a charge, and an `invoice.payment_failed` event will contain an invoice.  As with other API resources, you can use endpoints to retrieve an [individual event](https://stripe.com/docs/api#retrieve_event) or a [list of events](https://stripe.com/docs/api#list_events) from the API. We also have a separate [webhooks](http://en.wikipedia.org/wiki/Webhook) system for sending the `Event` objects directly to an endpoint on your server. Webhooks are managed in your [account settings](https://dashboard.stripe.com/account/webhooks), and our [Using Webhooks](https://stripe.com/docs/webhooks) guide will help you get set up.  When using [Connect](https://stripe.com/docs/connect), you can also receive notifications of events that occur in connected accounts. For these events, there will be an additional `account` attribute in the received `Event` object.  **NOTE:** Right now, access to events through the [Retrieve Event API](https://stripe.com/docs/api#retrieve_event) is guaranteed only for 30 days.
type Event struct {
	// The connected account that originated the event.
	Account *string `json:"account,omitempty"`
	// The Stripe API version used to render `data`. *Note: This property is populated only for events on or after October 31, 2014*.
	ApiVersion string `json:"api_version,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int                 `json:"created"`
	Data    NotificationEventData `json:"data"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Number of webhooks that have yet to be successfully delivered (i.e., to return a 20x response) to the URLs you've specified.
	PendingWebhooks int        `json:"pending_webhooks"`
	Request         NotificationEventRequest `json:"request,omitempty"`
	// Description of the event (e.g., `invoice.created` or `charge.refunded`).
	Type string `json:"type"`
}
