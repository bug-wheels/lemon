package discovery

import "lemon/cloud"

type DiscoveryClient interface {

	/**
	 * Gets all ServiceInstances associated with a particular serviceId.
	 * @param serviceId The serviceId to query.
	 * @return A List of ServiceInstance.
	 */
	GetInstances(serviceId string) ([]cloud.ServiceInstance, error)

	/**
	 * @return All known service IDs.
	 */
	GetServices() ([]string, error)
}
